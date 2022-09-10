package main

import "fmt"

func main() {

	// defer digunakan untuk mengakhiri eksekusi sebuah statement tepat sebelum blok fungsi selesai
	// selalu ingat bahwa defer digunakan untuk mengakhiri blok fungsi, bukan blok lainnya seperti blok seleksi dan kondis

	defer fmt.Println("arigato gozaimas")
	fmt.Println("selamat datang")
	orderSomeFood("Pizza")
	orderSomeFood("burger")
}

func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silakan tunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan Tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda: ", menu)
}
