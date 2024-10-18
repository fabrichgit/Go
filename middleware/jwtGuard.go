package middleware

import (
	"context"
	"net/http"
	"prime/helper"
	"strings"
)

// const claimsKey key = "user"

func JwtGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "", http.StatusServiceUnavailable)
			return
		}

		token = strings.Split(token, " ")[1]

		claims, err := helper.ValidateJWT(token)
		if err != nil {
			http.Error(w, "", http.StatusServiceUnavailable)
			return
		}
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// func PUser(r http.Request, claims any) {
// 	claims = r.Context().Value("user").(jwt.MapClaims)
// }
