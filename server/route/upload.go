package route

import (
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"get.cutie.cafe/rainy/conf"
	"get.cutie.cafe/rainy/upload"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
const numFilenameChars = 9

var rx *rand.Rand

func init() {
	rx = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func getExtension(file string) string {
	fx := strings.Split(file, ".")

	if len(fx) <= 1 {
		return ""
	}

	return strings.ReplaceAll(fx[len(fx)-1], "/", "")
}

func arrayContains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}

	return false
}

func isBlacklistedExtension(extension string) bool {
	return arrayContains(strings.Split(conf.GetString("BLACKLISTED_EXTENSIONS"), ","), extension)
}

func getRandomFilename(extension string) string {
	fn := ""

	for len(fn) < numFilenameChars {
		fn += string(chars[rx.Intn(len(chars))])
	}

	return fn + "." + extension
}

func checkPassword(password string) bool {
	pw := conf.GetString("UPLOAD_PASSWORD")

	return pw == "" || password == pw
}

func postUpload(uploader upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		if !checkPassword(r.Form.Get("password")) {
			w.WriteHeader(401)
			w.Write([]byte("Invalid password"))
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			log.Printf("error receiving file: %v", err)
			w.WriteHeader(400)
			w.Write([]byte("Error receiving file."))
			return
		}

		ext := getExtension(header.Filename)

		if isBlacklistedExtension(ext) {
			log.Printf("file was uploaded with blacklisted extension %s", ext)
			w.WriteHeader(400)
			w.Write([]byte("Error receiving file."))
			return
		}

		if uploader.MaxFileSize() < uint64(header.Size) {
			log.Printf("file size %d > max size %d", header.Size, uploader.MaxFileSize())
			w.WriteHeader(400)
			w.Write([]byte("File too big."))
			return
		}

		fn := ""

		for fn == "" || uploader.FileExists(fn) {
			fn = getRandomFilename(ext)
		}

		pf, err := uploader.StoreFileStream(fn, file)

		if err != nil {
			log.Printf("Error storing file %s: %v", fn, err)
			w.WriteHeader(500)
			w.Write([]byte("Error storing file. Please try again later."))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(*pf))

		file.Close()
	}
}

type SimpleModeVars struct {
	Meta     *InstanceMeta
	Password string
}

func getUpload(uploader upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		meta := genMeta(uploader)

		if meta.HasUploadPassword && conf.GetString("UPLOAD_PASSWORD") != r.Form.Get("password") {
			r.Form.Set("password", "")
		}

		err := runTemplate("html/boomer.html", SimpleModeVars{Meta: genMeta(uploader), Password: r.Form.Get("password")}, w)
		if err != nil {
			log.Printf("error parsing boomer template: %v", err)
			w.WriteHeader(500)
			w.Write([]byte("Oops, there's a bit of a problem."))
			return
		}
	}
}
