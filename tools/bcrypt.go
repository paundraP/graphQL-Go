package tools

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashed)
}

func CompareHashPassword(pw, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
