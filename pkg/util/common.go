package util

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

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

// NumToBHex 10进制数转换
// n 表示进制,16 or 36
func NumToBHex(number, n int) string {
	str := ""
	for number != 0 {
		remainder := number % n
		str = string(num2char[remainder]) + str
		number = number / n
	}
	return strings.ToUpper(str)
}

// BHex2Num 进制数转换
// n 表示进制,16 or 36
func BHex2Num(str string, n int) int {
	str = strings.ToLower(str)
	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒叙
	}
	return int(v)
}
