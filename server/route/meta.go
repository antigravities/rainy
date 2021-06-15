package route

import (
	"encoding/json"
	"net/http"

	"get.cutie.cafe/rainy/conf"
	"get.cutie.cafe/rainy/upload"
)

type InstanceMeta struct {
	Name               string  `json:"instanceName"`
	MaxUploadSizeMB    float32 `json:"maxUploadSizeMB"`
	ContactAbuse       string  `json:"contactAbuse"`
	ExtensionBlacklist string  `json:"extensionBlacklist"`
	HasUploadPassword  bool    `json:"hasUploadPassword"`
	Tagline            string  `json:"tagline"`
}

var metaJson *[]byte
var meta *InstanceMeta

func genMeta(u upload.Uploader) *InstanceMeta {
	if meta == nil {
		meta = &InstanceMeta{
			Name:               conf.GetString("INSTANCE_NAME"),
			MaxUploadSizeMB:    float32(u.MaxFileSize() / 1000000),
			ContactAbuse:       conf.GetString("CONTACT_ABUSE"),
			ExtensionBlacklist: conf.GetString("BLACKLISTED_EXTENSIONS"),
			HasUploadPassword:  conf.GetString("UPLOAD_PASSWORD") != "",
			Tagline:            conf.GetString("TAGLINE"),
		}
	}

	return meta
}

func getMeta(u upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if metaJson == nil {
			mb, err := json.Marshal(genMeta(u))

			if err != nil {
				mb = nil
				w.WriteHeader(500)
				w.Write([]byte("Try again later."))
				return
			}

			metaJson = &mb
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write(*metaJson)
	}
}

func postMeta(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if checkPassword(r.Form.Get("password")) {
		w.WriteHeader(204)
		w.Write([]byte{})
		return
	}

	w.WriteHeader(401)
	w.Write([]byte("Invalid password"))
}
