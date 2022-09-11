package main

import (
	"encoding/json"
	"fmt"
)

type User struct { // struct user akan digunakan untuk membuat variabel penumpang hasil decode json string
	FullName string `json:"Name"`
	Age      int
}

func main() {
	// Sumber data bisa berupa variabel objek cetakan struct, map[string]interface{}, atau slice.
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
