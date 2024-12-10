package common

import (
	"github.com/google/uuid"
)

func GenerateGuid() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
