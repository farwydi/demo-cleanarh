package domain

import "golang.org/x/crypto/bcrypt"

func PasswordTakeHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func PasswordCompare(pwdHash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(pwdHash, []byte(password))
}
