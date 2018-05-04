package gotool

import (
	"math/rand"
	"time"
)

var allChars = "0123456789abcdefghijklmnopqrstuvwxyz"
var allNumChars = "0123456789"

func MakeRandomString(l int) string {
	return MakeRandomStringFromTemplate(l, allChars)
}

func MakeRandomNumString(l int) string {
	return MakeRandomStringFromTemplate(l, allNumChars)
}

func MakeRandomStringFromTemplate(l int, template string) string {
	str := template
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
