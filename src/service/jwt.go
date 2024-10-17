package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Clé secrète pour signer le token
var jwtKey = []byte("ma_clé_secrète")

// Structure des revendications du JWT (claims)
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Fonction pour générer un JWT
func GenerateJWT(username string) string {
	// Définir la durée d'expiration du token
	expirationTime := time.Now().Add(5 * time.Minute)

	// Créer les claims avec un utilisateur et une durée d'expiration
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Créer le token avec la méthode de signature et les claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signer le token avec la clé secrète
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
