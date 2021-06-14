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
}

var meta *[]byte

func getMeta(u upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if meta == nil {
			mb, err := json.Marshal(&InstanceMeta{
				Name:               conf.GetString("INSTANCE_NAME"),
				MaxUploadSizeMB:    float32(u.MaxFileSize() / 1000000),
				ContactAbuse:       conf.GetString("CONTACT_ABUSE"),
				ExtensionBlacklist: conf.GetString("BLACKLISTED_EXTENSIONS"),
			})

			if err != nil {
				meta = nil
				w.WriteHeader(500)
				w.Write([]byte("Try again later."))
				return
			}

			meta = &mb
		}

		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write(*meta)
	}
}
