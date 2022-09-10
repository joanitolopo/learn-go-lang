package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//  encodetostring dan decodestring
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded: ", encodedString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodeByte)
	fmt.Println("decode: ", decodedString)
}
