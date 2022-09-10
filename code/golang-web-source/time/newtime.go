package main

import (
	"fmt"
	"time"
)

func main() {
	// setelah jeda waktu yang ditentukan sebuah data akan dikirimkan lewat channel

	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	// time.AfterFunc() --> dieksekusi jika waktusudah memenuhi durasi timer

	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("fungsi time.AfterFunc(). start...")
	<-ch
	fmt.Println("finish")

}
