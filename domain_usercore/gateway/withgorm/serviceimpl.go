package withgorm

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type commonImplementation struct {
}

func (r *commonImplementation) HashAndSaltPassword(ctx context.Context, password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (r *commonImplementation) ComparePassword(ctx context.Context, hashPass string, plainPass []byte) bool {
	byteHash := []byte(hashPass)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
