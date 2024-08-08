package kit

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

var (
	Uppercase       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase       = "abcdefghijklmnopqrstuvwxyz"
	Numbers         = "0123456789"
	Symbols         = ".,;:!?./-\"'#{([-|\\@)]=}*+"
	DefaultAlphabet = Uppercase + Lowercase + Numbers
	DefaultLength   = 64
)

func AllAlphabet(s string) (alphabet string) {
	if strings.ContainsAny(s, "u") {
		alphabet += Uppercase
	}
	if strings.ContainsAny(s, "l") {
		alphabet += Lowercase
	}
	if strings.ContainsAny(s, "n") {
		alphabet += Numbers
	}
	if strings.ContainsAny(s, "s") {
		alphabet += Symbols
	}
	return alphabet
}

func ShuffleArray(alphabet string, length int) string {
	if len(alphabet) == 0 {
		alphabet = DefaultAlphabet
	}
	if length < 1 || length > 512 {
		length = DefaultLength
	}
	s := strings.Repeat(alphabet, length)
	a := strings.Split(s, "")
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i)
		a[i], a[j] = a[j], a[i]
	}
	return strings.Join(a, "")[0:length]
}

func GenerateMD5(hashString, hashEncoding string) string {
	hasher := md5.New()
	// 将数据写入哈希对象
	hasher.Write([]byte(hashString))
	// 计算哈希值
	hashBytes := hasher.Sum(nil)
	return convertHash(hashBytes, hashEncoding)
}

func GenerateSha1(hashString, hashEncoding string) string {
	hasher := sha1.New()
	// 将数据写入哈希对象
	hasher.Write([]byte(hashString))
	// 计算哈希值
	hashBytes := hasher.Sum(nil)
	return convertHash(hashBytes, hashEncoding)
}

func convertHash(hashBytes []byte, hashEncoding string) string {
	switch hashEncoding {
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
