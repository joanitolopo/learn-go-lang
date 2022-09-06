package main

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type hitung interface { // embedded interfaces
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) diameter() float64 {
	return k.sisi * 2
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunRuang hitung = &kubus{4} // variabel objek struct yang menggunakan method pointer dan ditampung di variabel interface harus diambil referensinya terlebih dahulu

	fmt.Println("===== kubus")
	fmt.Println("luas \t\t: ", bangunRuang.luas())
	fmt.Println("keliling \t: ", bangunRuang.keliling())
	fmt.Println("keliling \t: ", bangunRuang.volume())
	fmt.Println("sisi \t\t: ", bangunRuang.(*kubus).sisi)

}
