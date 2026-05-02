package middleware

import (
	"context"
	jwtutil "market/internal/adapters/jwt"
	"net/http"
)

func AuthMiddleware(jwtM *jwtutil.Manager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "no token", http.StatusUnauthorized)
				return
			}

			const prefix = "Bearer "
			if len(authHeader) < len(prefix) {
				http.Error(w, "NO token", http.StatusUnauthorized)
				return
			}

			tokenStr := authHeader[len(prefix):]

			claims, err := jwtM.Parse(tokenStr)
			if err != nil {
				http.Error(w, "CANT PARSE", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "role", string(claims.Role))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
