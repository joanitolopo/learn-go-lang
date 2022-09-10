package main

import (
	"fmt"
	"strconv"
)

func main() {
	// konversi string ke int
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Printf("124 tipe %T menjadi 124 tipe %T\n", str, num) // 124 tipe string menjadi 124 tipe int

	}

	// konversi int ke string
	var num1 = 124
	var str2 = strconv.Itoa(num)
	fmt.Printf("124 tipe %T menjadi 124 tipe %T\n", num1, str2) // 124 tipe int menjadi 124 tipe string

	// konversi string berbentuk numerik dengan basis ke tipe numerik non-desimal dengan lebar data bisa ditentukan
	var str3 = "1010"
	var num3, err3 = strconv.ParseInt(str3, 2, 8) // basis 2 dan int8

	if err3 == nil {
		fmt.Printf("1010 tipe %T menjadi %v tipe %T (basis 2 tipe int8)\n", str3, num3, num3) // 1010 tipe string menjadi 10 tipe int64 (basis 2 tipe int8)
	}

	// konversi data numerik int64 ke string dengan basis numerik bisa tentuka sendiri
	var num4 = int64(24)
	var str4 = strconv.FormatInt(num4, 8)

	fmt.Printf("24 tipe %T menjadi %v tipe %T (basis 8)\n", num4, str4, str4) // 24 tipe int64 menjadi 30 tipe string (basis 8)

	// konversi string ke numerik desimal dengan lebar data bisa ditentukan
	var str5 = "24.12"
	var num5, err5 = strconv.ParseFloat(str5, 32)

	if err5 == nil {
		fmt.Printf("24.12 tipe %T menjadi %v tipe %T\n", str5, num5, num5) // 24.12 tipe string menjadi 24.1200008392334 tipe float64
	}

	// konversi float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data
	var num6 = float64(24.12)
	var str6 = strconv.FormatFloat(num6, 'f', 6, 64)

	fmt.Printf("24.12 tipe %T menjadi %v tipe %T\n", num6, str6, str6) // 24.12 tipe float64 menjadi 24.120000 tipe string

	// konversi string ke bool
	var str7 = "true"
	var bul, err7 = strconv.ParseBool(str7)

	if err7 == nil {
		fmt.Printf("true tipe %T menjadi %v tipe %T\n", str7, bul, bul) // true tipe string menjadi true tipe bool
	}

}
