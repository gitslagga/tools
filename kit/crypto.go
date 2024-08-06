package kit

import (
	"math/rand"
	"strings"
)

var (
	Uppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase     = "abcdefghijklmnopqrstuvwxyz"
	Numbers       = "0123456789"
	Symbols       = ".,;:!?./-\"'#{([-|\\@)]=}*+"
	DefaultStr    = Uppercase + Lowercase + Numbers
	DefaultLength = int64(64)
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
	if len(str) == 0 {
		str = DefaultStr
	}
	return str
}

func ShuffleArray(str string, length int64) string {
	if length < 1 || length > 512 {
		length = DefaultLength
	}
	s := strings.Repeat(str, int(length))
	a := strings.Split(s, "")
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i)
		a[i], a[j] = a[j], a[i]
	}
	return strings.Join(a, "")[0:length]
}
