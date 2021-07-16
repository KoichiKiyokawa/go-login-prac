package util

import "golang.org/x/crypto/bcrypt"

func HashString(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashAndString(hash string, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
