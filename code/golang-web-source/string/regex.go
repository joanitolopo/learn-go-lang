package main

import (
	"fmt"
	"regexp"
)

func main() {
	// regex pencarian karakter
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	// mencari semua string yang sesuai dengan ekspresi regex, dengan kembalian berupa slice string
	var res1 = regex.FindAllString(text, 2) // mengembalikan max 2 data saja
	fmt.Printf("%#v \n", res1)              // []string{"banana", "burger"}
	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2) // []string{"banana", "burger", "soup"}

	// mendeteksi apakah string memenuhi pola regexp
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch) // true

	// mencari string yang memenuhi kriteria regexp yang ditentukan
	var str = regex.FindString(text)
	fmt.Println(str) // banana (jika ada banyak substring yang sesuai, maka dikembalikan yang pertama saja)

	// mencari index string kembalian hasil dari operasi
	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)       // [0 6]
	fmt.Println(text[0:6]) // banana

	// mencari banyak string yang memenuhi
	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1) // [banana burger soup]
	var str2 = regex.FindAllString(text, 2)
	fmt.Println(str2) // [banana burger]

	// replace semua string yang memenuhi regexp
	var str3 = regex.ReplaceAllString(text, "potato")
	fmt.Println(str3) // potato potato potato

	// digunakna untuk replace semua, dengan kondisi yang bisa ditentukan untuk setiap substring
	var str4 = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})

	fmt.Println(str4) // banana potato soup

}
