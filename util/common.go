package util

import (
	"encoding/base64"
	"math/rand"
	"time"
)

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func EncodeStringToBase64(s string) string {
	// 将字符串转换为字节数组
	data := []byte(s)
	// 将字节数组进行base64编码
	encoded := base64.StdEncoding.EncodeToString(data)
	// 返回编码后的字符串
	return encoded
}

func DecodeBase64ToString(encoded string) (string, error) {
	// 将base64编码的字符串转换为字节数组
	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	// 将字节数组转换为字符串
	decoded := string(data)
	// 返回解码后的字符串
	return decoded, nil
}
