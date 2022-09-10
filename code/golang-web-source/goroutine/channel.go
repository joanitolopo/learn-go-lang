package main

import (
	"fmt"
	"runtime"
)

func main() {

	// channel merupakan variabel yang menjadi pengirim atau penerima data di sebuah goroutine

	runtime.GOMAXPROCS(2)

	var messages = make(chan string) // pembuatan channel, variabel channel memiliki satu tugas, menjadi pengirim atau penerima data
	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data // variabel data dikirim melalui channel messages
	}

	// goroutine dilakukan secara asynchronous sedangkan serah-terima data channel adalah synchronous
	// jadi, goroutine yang selesai paling awal akan mengirim data lebih dulu

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	// dari ketiga fungsi diatas, yang selesai paling awal akan mengirim data lebih dulu, yang datanya
	// kemudian diterima variabel message1

	// penerimaan channel bersifat blocking, artinya statement setelah var message 1,tidak akan diekesui sampai ada data yang dikirim lewat channel
	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)

	// output:
	// hello john wick
	// hello jason bourne
	// hello ethan hunt
}
