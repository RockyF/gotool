package gotool

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomGetFromArray(array []interface{}) interface{} {
	return array[random.Intn(len(array))]
}
