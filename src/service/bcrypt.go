package service

import (
	"golang.org/x/crypto/bcrypt"
)

// Fonction pour hacher un mot de passe
func hashPassword(password string) string {
	// Le deuxième argument est le coût (recommandé : bcrypt.DefaultCost)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // renvoie vrai si les deux correspondent
}
