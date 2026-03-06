package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router http.Handler
}

func NewServer(router *chi.Mux) *Server {
	return &Server{
		router: router,
	}
}

func (srv *Server) Start() {
	fmt.Println("Server is listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", srv.router)
	if err != nil {
		return
	}
}
