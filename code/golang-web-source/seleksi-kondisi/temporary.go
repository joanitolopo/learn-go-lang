package main

import "fmt"

func main() {
	var point = 8840.0

	// Variabel percent nilainya didapat dari hasil perhitungan,
	// dan hanya bisa digunakan di deretan blok seleksi kondisi itu saja.
	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
