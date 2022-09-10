package main

import (
	"fmt"
	"os"
)

func main() {
	// fungsi os.exitakan memaksa program berhenti secara paksa, semua statement setelah exit tidak akan dijalankan termasuk defer

	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")

	// output:
	// exit status 1
}
