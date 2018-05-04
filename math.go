package gotool

import (
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomGetFromArray(array interface{}) interface{} {
	temp := array.([]interface{})

	return temp[random.Intn(len(temp))]
}
