package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	// var lulusUjian bool = ujian >= 80
	// fmt.Println("Lulus Ujian is ", lulusUjian)
	// var lulusAbsensi bool = absensi >= 80
	// fmt.Println("Lulus absensi is ", lulusAbsensi)

	// var lulus bool = lulusUjian && lulusAbsensi

	// fmt.Println("Anda dinyataka: ", lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

}
