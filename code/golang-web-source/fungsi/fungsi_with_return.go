package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Sebuah fungsi bisa dirancang tidak mengembalikan nilai balik (void), atau bisa mengembalikan suatu nilai.
	// Fungsi yang memiliki nilai kembalian, harus ditentukan tipe data nilai baliknya pada saat deklarasi.

	rand.Seed(time.Now().Unix()) // Fungsi ini diperlukan untuk memastikan bahwa angka random yang akan di-generate benar-benar acak
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)

}

func randomWithRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value

}
