package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"ims-intro/models"
	"net/http"
	"strings"
)

var JwtKey = []byte("your_secret_key")

func AuthAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			http.Error(w, "No token provided", http.StatusUnauthorized)
			return
		}

		tokenStr = strings.Split(tokenStr, "Bearer ")[1]
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid || claims.Role != "admin" {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
