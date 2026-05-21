package internal

import (
	"crypto/sha1"
	"fmt"
)

func TreeHash(object string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(object))
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash, nil
}
