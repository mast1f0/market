package http

import (
	"fmt"
	"net/http"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (srv *Server) Start() {
	fmt.Println("Server is listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", srv.handlers.Router)
	if err != nil {
		return
	}
}
