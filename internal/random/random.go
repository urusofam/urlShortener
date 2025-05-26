package random

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rnd.Intn(len(letterBytes))]
	}

	return string(b)
}
