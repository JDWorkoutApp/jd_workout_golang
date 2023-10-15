package helper

import (
	"math/rand"
	"time"
)

func RandString(length int) string {
	letters := []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}