package main

import (
	"fmt"
	"reflect"
)

func main() {

	var number = 23
	var reflectValue = reflect.ValueOf(number) // menyimpan informasi mengenai variabel number

	fmt.Println("tipe variabel: ", reflectValue.Type()) // mengembalikan informasi tipe datanya

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel: ", reflectValue.Int())
	}

	// Pengaksesan nilai dalam bentuk interface{}
	fmt.Println("nilai variabel using interface: ", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	fmt.Println("nilai: ", nilai)

}
