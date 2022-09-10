package main

import (
	"fmt"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func wathcer(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\ntime out ! no answer more than", timeout, "second")
	os.Exit(0)
}

func main() {
	var timeout = 5
	var ch = make(chan bool)

	go timer(timeout, ch)
	go wathcer(timeout, ch)

	var input string
	fmt.Print("What is 725/25 : ")
	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong")
	}

	// ketika user tidak input apa-apa selama 5 detik,maka akan muncul pesan timeout
}
