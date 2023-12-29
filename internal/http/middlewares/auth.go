package middlewares

import (
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth"
)

var ja *jwtauth.JWTAuth

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		a := strings.Split(authHeader, "Bearer ")

		if len(a) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid Bearer token"}`))
			return
		}

		jwtToken := a[1]

		tokenDecoded, err := ja.Decode(jwtToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"message": "Invalid Bearer token"}`))
			return
		}

		// token
		next.ServeHTTP(w, r)
	})
}
