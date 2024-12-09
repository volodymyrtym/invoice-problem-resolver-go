package idgenerator

import (
	"github.com/google/uuid"
)

func GenerateUserID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
