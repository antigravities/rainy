package route

import (
	"io/fs"
	"net/http"

	"get.cutie.cafe/rainy/conf"
	"get.cutie.cafe/rainy/static"
	"get.cutie.cafe/rainy/upload"
	"github.com/gorilla/mux"

	"html/template"
)

type Router struct {
	Uploader upload.Uploader
	Router   *mux.Router
}

// New instantiates and binds routes to a new Router and *mux.Router.
func New(uploader upload.Uploader) *Router {
	router := mux.NewRouter()

	subfs, err := fs.Sub(static.ModernContent, "html/modern")
	if err != nil {
		panic(err)
	}

	router.HandleFunc("/upload", postUpload(uploader)).Methods("POST")
	router.HandleFunc("/upload", getUpload(uploader)).Methods("GET")
	router.HandleFunc("/meta", getMeta(uploader)).Methods("GET")
	router.HandleFunc("/meta", postMeta).Methods("POST")
	router.HandleFunc("/", getIndex(uploader))

	if conf.GetInt("SHOULD_SERVE") == 1 {
		// TODO: http.FileServer()
		router.HandleFunc("/f/{file}", getFile(uploader))
	}

	router.PathPrefix("/").Handler(http.FileServer(http.FS(subfs)))

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
