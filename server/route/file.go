package route

import (
	"log"
	"mime"
	"net/http"

	"get.cutie.cafe/rainy/upload"
	"github.com/gorilla/mux"
)

func getFile(uploader upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if val, ok := vars["file"]; ok && uploader.FileExists(val) {
			bx, err := uploader.GetFile(val)

			if err != nil {
				log.Printf("error reading %s: %v", val, err)
				w.WriteHeader(500)
				w.Write([]byte("Could not read file right now, try again later"))
				return
			}

			mimetype := mime.TypeByExtension(getExtension(val))

			if mimetype == "" {
				mimetype = "application/octet-stream"
			}

			w.WriteHeader(200)
			w.Header().Add("Content-Type", mimetype)

			w.Write(bx)

			return
		}

		w.WriteHeader(404)
		w.Write([]byte("Not found"))
	}
}
