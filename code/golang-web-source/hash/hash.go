package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	// hash adalah algoritma enkripsi untuk mengubah text menjadi deretan angka acak
	var text = "this is secret"
	var sha = sha1.New()         // objek bertipe hash.Hash
	sha.Write([]byte(text))      // set data yang akan dihash, harus dalam bentuk []byte
	var encrypted = sha.Sum(nil) // digunakan untuk eksekusi proses hash, method ini memutuhkan paramater, isi dengan nil
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString) // f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8

}
