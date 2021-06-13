package route

import (
	"net/http"

	"get.cutie.cafe/rainy/static"
	"github.com/gorilla/mux"

	"text/template"
)

func New() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", getIndex)

	return router
}

func runTemplate(templateFile string, obj interface{}, w http.ResponseWriter) error {
	file, err := static.Get(templateFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New(templateFile).Parse(string(file))
	if err != nil {
		return err
	}

	err = tmpl.Execute(w, obj)
	if err != nil {
		return err
	}

	return nil
}
