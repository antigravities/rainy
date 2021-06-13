package server

import (
	"log"
	"net/http"

	"get.cutie.cafe/rainy/server/route"
	"get.cutie.cafe/rainy/upload"
)

type Server struct {
	uploader *upload.Uploader
}

func New(uploader *upload.Uploader) *Server {
	return &Server{
		uploader: uploader,
	}
}

func (s *Server) Listen(addr string) {
	http.Handle("/", route.New())

	log.Printf("Preparing to listen on %s", addr)
	http.ListenAndServe(addr, nil)
}
