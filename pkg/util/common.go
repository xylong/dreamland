package util

import (
	"math/rand"
	"time"
)

// RandomString 返回随机字符串
func RandomString(n uint8) string {
	letters := []byte("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
