package image

import (
	"fmt"
	"math/rand"

	"api-server/config"
)

func newRandomImagePath() (string, string) {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 64)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s/%s", config.ImageDirectory, string(b)), string(b)
}
