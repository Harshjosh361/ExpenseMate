package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	//default cost to balance performance and security
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func Checkpassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
