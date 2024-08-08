package kit

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

var (
	Uppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase     = "abcdefghijklmnopqrstuvwxyz"
	Numbers       = "0123456789"
	Symbols       = ".,;:!?./-\"'#{([-|\\@)]=}*+"
	DefaultStr    = Uppercase + Lowercase + Numbers
	DefaultLength = 64
)

func AllAlphabet(s string) (str string) {
	if strings.ContainsAny(s, "u") {
		str += Uppercase
	}
	if strings.ContainsAny(s, "l") {
		str += Lowercase
	}
	if strings.ContainsAny(s, "n") {
		str += Numbers
	}
	if strings.ContainsAny(s, "s") {
		str += Symbols
	}
	return str
}

func ShuffleArray(str string, length int) string {
	if len(str) == 0 {
		str = DefaultStr
	}
	if length < 1 || length > 512 {
		length = DefaultLength
	}
	s := strings.Repeat(str, length)
	a := strings.Split(s, "")
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i)
		a[i], a[j] = a[j], a[i]
	}
	return strings.Join(a, "")[0:length]
}

func GenerateMD5(s, e string) string {
	hasher := md5.New()
	// 将数据写入哈希对象
	hasher.Write([]byte(s))
	// 计算哈希值
	hashBytes := hasher.Sum(nil)

	switch e {
	case "b":
		// 将哈希值转换为二进制字符串
		b := fmt.Sprintf("%08b", hashBytes)
		b = strings.ReplaceAll(b[1:len(b)-1], " ", "")
		return b
	case "x":
		// 将哈希值转换为十六进制字符串
		x := fmt.Sprintf("%x", hashBytes)
		return x
	case "b64":
		// 将哈希值转换为Base64字符串
		b64 := base64.StdEncoding.EncodeToString(hashBytes)
		return b64
	case "b64u":
		// 将哈希值转换为Base64url字符串
		b64u := base64.RawURLEncoding.EncodeToString(hashBytes)
		return b64u
	}

	return hex.EncodeToString(hashBytes)
}
