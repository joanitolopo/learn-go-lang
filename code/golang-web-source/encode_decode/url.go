package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "https://github.com/joanitolopo"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
