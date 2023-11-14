package helper

import "golang.org/x/crypto/bcrypt"

func HashThePass(pwd string) string {
	salt := 8
	password := []byte(pwd)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func CompareThePass(hashing, password []byte) bool {
	hash, pass := []byte(hashing), []byte(password)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}