package main

import "fmt"

func main() {
	// Di perulangan bersarang, break dan continue akan berlaku pada blok perulangan di mana ia digunakan saja.
	// Ada cara agar kedua keyword ini bisa tertuju pada perulangan terluar atau perulangan tertentu, yaitu dengan memanfaatkan teknik pemberian label.
outerLoop:
	for i := 0; 1 < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("Matriks [", i, "][", j, "]", "\n")
		}
	}
}
