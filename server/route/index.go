package route

import (
	"log"
	"net/http"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	err := runTemplate("html/index.html", struct{}{}, w)
	if err != nil {
		log.Println(err)
	}
}
