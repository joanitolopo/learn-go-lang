package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64 // definisi method dalam interface
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

type persegi struct {
	sisi float64
}

// lingkaran
func (l lingkaran) jariJari() float64 { // function jari-jari yang mempunyai method lingkaran
	return l.diameter / 2
}

func (l lingkaran) luas() float64 { // function luas yang mempunyai method lingkaran
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 { // function keliling yang mempunyai method lingkaran
	return math.Pi * l.diameter

}

// persegi
func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	// Interface adalah kumpulan definisi method yang tidak memiliki isi (hanya definisi saja)
	// interface baru bisa digunakan jika sudah ada isinya

	var bangunDatar hitung // variabel objek bangunDatar yang bertipe interface hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas\t\t: ", bangunDatar.luas())       // 100
	fmt.Println("keliling\t: ", bangunDatar.keliling()) // 40

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas\t\t: ", bangunDatar.luas())                    // 153.93
	fmt.Println("keliling\t: ", bangunDatar.keliling())              // 43.98
	fmt.Println("jari-jari\t: ", bangunDatar.(lingkaran).jariJari()) // 7

}
