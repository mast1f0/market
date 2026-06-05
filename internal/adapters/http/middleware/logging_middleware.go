package middleware

import (
	"log/slog"
	"market/internal/engine/logger"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(log *logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wr := &responseWriter{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			defer func() {
				duration := time.Since(start)

				args := []any{
					slog.String("method", r.Method),
					slog.String("path", r.URL.Path),
					slog.String("query", r.URL.RawQuery),
					slog.Int("status", wr.status),
					slog.String("duration", duration.String()),
					slog.String("user_agent", r.UserAgent()),
				}

				switch {
				case wr.status >= 500:
					log.Error("request failed", args...)
				case wr.status >= 400:
					log.Warn("client error", args...)
				default:
					log.Info("request completed", args...)
				}
			}()

			next.ServeHTTP(wr, r)
		})
	}
}
