package main

import "fmt"

func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")

		func() { // jika tidak dibungkus dengan fungsi, maka defer akan tetap di jalankan ketika blok func main berakhir
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}
