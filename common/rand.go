package common

import (
	"math/rand"
	"time"
)

const alphabet  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func GenerateRandomString(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[r.Int63() % int64(len(alphabet))]
	}
	return string(b)
}