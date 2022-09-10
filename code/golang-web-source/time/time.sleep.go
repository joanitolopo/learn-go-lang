package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 seconds")
	fmt.Println("===================")

	for true {
		fmt.Println("Hello !!")
		time.Sleep(1 * time.Second)
	} // schedulre untuk menampilkan pesan halo setiap 1 detik
}
