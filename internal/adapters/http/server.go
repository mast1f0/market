package http

import (
	"context"
	"log/slog"
	"market/internal/engine/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *logger.Logger
}

func NewServer(addr string, handler http.Handler, logger *logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		logger: logger,
	}
}

func (srv *Server) Start() error {
	go func() {
		srv.logger.Info("Server is listening on http://localhost", srv.httpServer.Addr)
		err := http.ListenAndServe(":8080", srv.httpServer.Handler)
		if err != nil {
			return
		}
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-shutdown:
		srv.logger.Info("Received signal: %v", slog.String("signal", sig.String()))

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.httpServer.Shutdown(ctx); err != nil {
			if closeErr := srv.httpServer.Close(); closeErr != nil {
				srv.logger.Info("Force close error: %v", closeErr)
			}
			srv.logger.Error("graceful shutdown failed:", err)
			return err
		}

		srv.logger.Info("Server stopped gracefully")
		return nil
	}
}
