package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano()) // salt yang digunakan adalah hasil ekspresi time.Now().UnixNano()
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	// salt adalah data acak yang digabungkan pada data asli sebelum proses hash dilakukan
	var text = "this is secret"
	fmt.Printf("original: %s\n\n", text)

	// hasil dari setiap hash nya akan selalu unik setiap detiknya karena scope terendah waktu pada fungsi adalah nano second
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1: %s\n\n", hashed1) // 53825f0d71253cd18f075594daa3383d87abc040

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2: %s\n\n", hashed2) // b38ff924c243ce4afc503beef74c6574e248a8c8

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3: %s\n\n", hashed3) // fbe3ce2e0c653f0cf40b38b1a9ea6f62a5d5d667

	_, _, _ = salt1, salt2, salt3

}
