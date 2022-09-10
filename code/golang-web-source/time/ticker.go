package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool) // variabel yang mengontrol kapan ticker harus berhenti
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second) // wait for 10 seconds
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}

	// selama 10 detik, di setiap detiknya akan muncul pesan halo
}
