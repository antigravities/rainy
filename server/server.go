package server

import (
	"net/http"

	"get.cutie.cafe/rainy/server/route"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Listen() {
	http.Handle("/", route.New())
	http.ListenAndServe("0.0.0.0:4000", nil)
}
