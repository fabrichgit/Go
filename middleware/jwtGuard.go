package middleware

import (
	"context"
	"net/http"
	"prime/helper"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JwtGuard(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return helper.JwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "jwtClaims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
