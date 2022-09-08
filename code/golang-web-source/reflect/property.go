package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func main() {

	var s1 = &student{Name: "wick", Grade: 2}
	s1.getPropertyInfo()
}

func (s *student) getPropertyInfo() { // function ini memiliki method ke student struct

	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr { // apakah objek tersebut pointer atau tidak
		reflectValue = reflectValue.Elem() // mengmabil objek reflect aslinya
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ { // perulangan sebanyak jumlah properti didalam struc
		fmt.Println("nama: ", reflectType.Field(i).Name)          // return name property
		fmt.Println("tipe data: ", reflectType.Field(i).Type)     // return tipe data
		fmt.Println("nilai: ", reflectValue.Field(i).Interface()) // return nilai property
		fmt.Println("")
	}
}
