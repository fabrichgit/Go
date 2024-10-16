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
	// Retourne le mot de passe haché sous forme de chaîne de caractères
	return string(hashedPassword)
}

// Fonction pour vérifier un mot de passe par rapport au hash
func checkPasswordHash(password, hash string) bool {
	// bcrypt.CompareHashAndPassword compare le mot de passe en clair avec le hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // renvoie vrai si les deux correspondent
}
