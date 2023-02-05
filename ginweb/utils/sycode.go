package utils

import (
	"math/rand"
	"time"
)

func SyCode() int32 {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := rnd.Int31n(1000000)
	return vcode

}
