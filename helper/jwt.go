package helper

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id string `json:"Id"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string) string {

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return ""
	}

	return tokenString
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func GetPayload(w http.ResponseWriter, r *http.Request) (string, error) {
	claims, ok := r.Context().Value("jwtClaims").(*jwt.MapClaims)
	if !ok {
		http.Error(w, "failed to get claims", http.StatusInternalServerError)
		return "", fmt.Errorf("failed to get claims")
	}

	id, ok := (*claims)["Id"].(string)
	if !ok {
		http.Error(w, "ID not found in token", http.StatusUnauthorized)
		return "", fmt.Errorf("ID not found in token")
	}

	return id, nil
}

// func GetPayload(r *http.Request) (string, error) {
// 	claims, ok := r.Context().Value(jwtmiddleware.ContextKey{}).(*jwt.MapClaims)
// 	if !ok {
// 		return "", fmt.Errorf("failed to get claims")
// 	}

// 	id, ok := (*claims)["Id"].(string)
// 	if !ok {
// 		return "", fmt.Errorf("ID not found in token")
// 	}

// 	return id, nil
// }
