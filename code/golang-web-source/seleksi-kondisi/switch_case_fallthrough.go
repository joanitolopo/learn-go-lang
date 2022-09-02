package main

import "fmt"

func main() {
	var point = 6

	// Keyword fallthrough digunakan untuk memaksa proses pengecekan diteruskan
	// ke satu case selanjutnya dengan tanpa menghiraukan nilai kondisinya,
	// jadi satu case di pengecekan selanjutnya tersebut selalu dianggap benar (meskipun aslinya adalah salah).

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Awesome")
		fallthrough
	case point < 5:
		fmt.Println("You need to learn more")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn more")
		}
	}
}
