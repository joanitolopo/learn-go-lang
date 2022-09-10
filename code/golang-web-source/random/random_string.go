package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := randomString(20)
	fmt.Println(a)
}

func randomString(length int) string {

	rand.Seed(time.Now().UTC().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ") //rune is alias fot int32

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)

}
