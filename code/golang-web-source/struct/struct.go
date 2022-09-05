package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

type mahasiswa struct {
	nama string
	nim  int
}

func main() {
	// Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method),
	// yang dibungkus sebagai tipe data baru dengan nama tertentu

	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name: ", s1.name)
	fmt.Println("grade: ", s1.grade)
	fmt.Println(strings.Repeat("=", 20))

	// 3 ways to declare object variable
	var st1 = mahasiswa{}
	st1.nama = "wick"
	st1.nim = 682019013

	var st2 = mahasiswa{"aldi", 672019013}

	st3 := mahasiswa{nama: "aldo"}

	st4 := mahasiswa{nama: "roy", nim: 56201854}

	st5 := mahasiswa{nim: 78202067, nama: "brucle"}

	fmt.Println("Student 1 :", st1.nama)
	fmt.Println("Student 2: ", st2.nama)
	fmt.Println("Student 3: ", st3.nama)
	fmt.Println("Student 4: ", st4.nama)
	fmt.Println("Student 5: ", st5.nama)

}
