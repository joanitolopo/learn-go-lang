package main

import "fmt"

func main() {
	var point = 8

	// Switch merupakan seleksi kondisi yang sifatnya fokus pada satu variabel,
	// lalu kemudian di-cek nilainya
	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

	// Perlu diketahui, switch pada pemrograman Go memiliki perbedaan dibanding bahasa lain.
	// Di Go, ketika sebuah case terpenuhi, tidak akan dilanjutkan ke pengecekan case selanjutnya,
	// meskipun tidak ada keyword break di situ. Konsep ini berkebalikan dengan switch pada umumnya,
	// yang ketika sebuah case terpenuhi, maka akan tetap dilanjut mengecek case selanjutnya kecuali ada keyword break
}
