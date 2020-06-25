package utils

import (
	"math/rand"
	"time"
)

// https://blog.csdn.net/impressionw/article/details/72765756

// RandomString xxx
func RandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	name := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		name[i] = letters[rand.Intn(len(letters))]
	}
	return string(name)
}
