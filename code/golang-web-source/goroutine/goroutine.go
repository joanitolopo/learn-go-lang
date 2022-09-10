package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {

	// goroutine adalah sebuah mini thread dengan kapastias 2kB untuk satu goroutine
	// goroutine bersifat asynchronous sehingga menjadikannya tidak saling tunggu dengan goroutine lainnya
	// goroutine berjalan dalam multi processor, kita bisa mennentuka berapa banyak core yang akfit

	runtime.GOMAXPROCS(2) // menentukan jumlah core untuk eksekusi program

	go print(5, "halo") // fungsi dijalankan sebagai goroutine baru
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)

	// --> output:
	// 1 apa kabar
	// 2 apa kabar
	// 3 apa kabar
	// 1 halo
	// 2 halo
	// 3 halo
	// 4 halo
	// 5 halo
	// 4 apa kabar
	// 5 apa kabar

}
