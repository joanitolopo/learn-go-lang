package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.contains digunakan untuk deteksi apakah string merupakan bagian dari string lainnya
	var isExist = strings.Contains("john wick", "wick")
	fmt.Println(isExist) // true

	//  deteksi apakah sebuah string diawali string tertentu
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true
	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	// deteksi apakah sebuah string diakhiri string tertentu
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false
	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	// menghitung jumlah karakter tertentu dari sebuah string
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	// mencari posisi indeks sebuah string dalam string
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2
	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4 (jika ada substring, maka yang pertama yang diambil)

	// mengganti bagian dari string dengan string tertentu
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1) // bonana (diganti hanya string pertama)
	fmt.Println(newText1)

	var newText2 = strings.Replace(text, find, replaceWith, 2) // bonona (diganti 2 string)
	fmt.Println(newText2)

	var newText3 = strings.Replace(text, find, replaceWith, -1) // bonono (mengganti semua)
	fmt.Println(newText3)

	// memisah string dengan tanda pemisah bisa ditentukan sendiri
	var string1 = strings.Split("the dark knight", " ") // output: ["the", "dark", "knight"]
	fmt.Println(string1)
	var string2 = strings.Split("batman", "") // output: ["b", "a", "t", "m", "a", "n"]
	fmt.Println(string2)

	// menggabungkan slice string menjadi sebuah string dengan pemisah tertentu
	var data = []string{"banana", "pepaya", "tomato"}
	var str = strings.Join(data, "-") // banana-pepaya-tomato
	fmt.Println(str)

	// mengubaha huruf-huruf string menjadi huruf kecil
	var str2 = strings.ToLower("aLAy")
	fmt.Println(str2) // alay
}
