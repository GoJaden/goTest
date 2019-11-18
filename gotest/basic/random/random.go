package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func GetRandString(strLen int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, strLen)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := -1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(GetRandString(32))
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(10))
	}
}

//方式一：获取随机多少位的字符串
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	//[]rune 可以用[]byte代替，节省内存开销
	b := make([]rune, n)
	for poi, _ := range b {
		b[poi] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
