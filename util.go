package ito

import (
	"crypto/sha256"
	"fmt"

	"github.com/denisbrodbeck/machineid"
)

func GenerateNodeID(appPrefix string) (string, error) {
	id, err := machineid.ID()
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(appPrefix + id))
	return fmt.Sprintf("%x", hash)[:16], nil
}
