package main

import (
	"fmt"
	"runtime"
)

func main() {

	// secara default, proses transfer data pada channel menggunakan cara un-buffered, atau tidak buffer di memori
	// hal ini menyebabkan proses serah terimanya bersifat blocking.
	// kita bisa membuat proses serah terimanya tidak bersifat blocking (asynchronous) dengan cara membuat buffer
	// selama tidak melebih jumlah buffer, proses pengiriman akan berjalan tidak blocking

	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3) // message dengan buffer, buffer dimulai dari 0 sehingga nilai buffernya maksimal 4

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}

	// output
	// send data 0
	// send data 1
	// send data 2
	// send data 3
	// receive data 0
	// receive data 1
	// receive data 2
	// receive data 3
	// send data 4
}
