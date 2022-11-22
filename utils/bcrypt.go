package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	bytePass := []byte(pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	pass = string(hPass)
	return pass
}

func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}
