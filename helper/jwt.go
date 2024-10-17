package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

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

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return ""
	}

	return tokenString
}

// Fonction pour vérifier et décoder un JWT
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Décoder et valider le token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Si le token n'est pas valide ou est expiré, renvoyer une erreur
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
