package main

import "fmt"

func main() {
	// Pointer adalah reference atau alamat memori. Variabel pointer berarti variabel yang berisi alamat memori suatu nilai.
	// Sebagai contoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori di mana nilai 4 disimpan, bukan nilai 4 itu sendiri.

	// ada dua hal penting:
	// Variabel biasa bisa diambil nilai pointernya (add '&' sebelum nama variabel)
	// nilai asli variabel pointer juga bisa diambil (add '*' sebelum nama variabel)

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("NumberA (value)\t\t :", numberA)  // 4
	fmt.Println("NumberA (address)\t :", &numberA) // 0xc0000b4000

	fmt.Println("NumberB (value)\t\t :", *numberB) //4
	fmt.Println("NumberB (address)\t :", numberB)  // 0xc0000b4000

}
