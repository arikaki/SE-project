package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"kora.com/project/src/database"
)

var SecretKey = []byte(os.Getenv("Cookie_Key"))

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("Session")
		tokenStr := cookie.Value
		claims := &database.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// check token signing method etc
			return SecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "Bad Request", http.StatusForbidden)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		foundUser, err := database.GetUser(claims.Username)
		if err != nil {
			http.Error(w, "User Not Found", http.StatusForbidden)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), 0, foundUser)
		rWithUser := r.WithContext(ctxWithUser)

		if err != nil {
			return
		}
		next.ServeHTTP(w, rWithUser)
	})
}
