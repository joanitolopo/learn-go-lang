package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	diameter := float64(15)
	tinggi := float64(66)
	luas_lingkaran, keliling_lingkaran := calculate(diameter)
	luas_tabung, keliling_tabung := calculated(diameter, tinggi)

	fmt.Printf("Luas lingkaran\t\t: %.2f \n", luas_lingkaran)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", keliling_lingkaran)
	fmt.Println(strings.Repeat("=", 50))

	fmt.Printf("Luas Tabung\t\t: %.2f \n", luas_tabung)
	fmt.Printf("Keliling Tabung\t\t: %.2f \n", keliling_tabung)

}

func calculate(d float64) (float64, float64) {
	// function with no predefined return value

	// hitung luas
	area := math.Pi * math.Pow(d/2, 2)
	// hitung keliling
	circumference := math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}

func calculated(d float64, t float64) (area float64, circumference float64) {
	// function with predefined return value
	area = 2 * math.Pi * d / 2
	circumference = math.Pi * math.Pow(d/2, 2) * t
	return

}
