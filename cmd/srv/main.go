package main

import (
	"market/internal/adapters/storage/memory"
	"market/internal/http"
)

func main() {
	srv := http.NewServer(http.NewHandlers(memory.NewMemory()))
	srv.Start()
}
