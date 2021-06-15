package route

import (
	"log"
	"net/http"

	"get.cutie.cafe/rainy/upload"
)

func getIndex(uploader upload.Uploader) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := runTemplate("html/modern/index.html", genMeta(uploader), w)
		if err != nil {
			log.Println(err)
		}
	}
}
