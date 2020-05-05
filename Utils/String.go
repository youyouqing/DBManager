package Utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func RandIntString(num int) string {
	rand.Seed(time.Now().UnixNano())
	randStr := ""
	for i := 0; i < num; i++ {
		randStr = strconv.Itoa(rand.Intn(10)) + randStr
	}
	return randStr
}
