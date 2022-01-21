package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	//Set random  seed
	rand.Seed(time.Now().Unix())
}

//Generate random function (n indicates how many bits are generated)
func MkRandom(n int) string {
	chars := `abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ0123456789`

	var builder strings.Builder
	for i := 0; i < n; i++{
	builder.WriteByte(chars[rand.Intn(len(chars))])
	}

	return builder.String()
}

func main() {
	passwd := MkRandom(6)
	fmt.Println(passwd)
}