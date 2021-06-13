package route

import (
	"net/http"

	"get.cutie.cafe/rainy/static"
	"get.cutie.cafe/rainy/upload"
	"github.com/gorilla/mux"

	"text/template"
)

type Router struct {
	Uploader upload.Uploader
	Router   *mux.Router
}

// New instantiates and binds routes to a new Router and *mux.Router.
func New(uploader upload.Uploader) *Router {
	router := mux.NewRouter()

	router.HandleFunc("/", getIndex)
	router.HandleFunc("/upload", postUpload(uploader))

	return &Router{
		Uploader: uploader,
		Router:   router,
	}
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
