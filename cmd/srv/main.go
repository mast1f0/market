package main

import (
	"market/internal/adapters/storage/memory"
	"market/internal/http"
	"market/internal/http/handlers"
	"market/internal/service"
)

func main() {

	repo := handlers.NewProductHandler(memory.NewMemory())
	srv := http.NewServer(http.NewHandlers(repo))
	srv.Start()
}
