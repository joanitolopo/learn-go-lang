package main

import (
	"fmt"
	"strings"
)

func main() {

	fruits := []string{"apple", "grape", "banana", "melon"} // if no memory initilisation then it is slice

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = fruits[1:2]
	var bbFruits = fruits[0:1]

	// You can try the %v, %+v or %#v verbs when using printf
	fmt.Printf("Slice Parents (fruits) %v: \n", fruits)
	fmt.Printf("Slice A: %v \n", aFruits)
	fmt.Printf("Slice B: %+v \n", bFruits)
	fmt.Printf("Slice a: %+v \n", aaFruits)
	fmt.Printf("Slice b: %v \n", bbFruits)

	// how if 'grape' value we change to 'pineaple
	fmt.Println(strings.Repeat("=", 55))
	bbFruits[0] = "pinneaple"

	fmt.Printf("Slice Parents after change(fruits): %v\n", fruits)
	fmt.Printf("Slice A after change: %v \n", aFruits)
	fmt.Printf("Slice B after change: %+v \n", bFruits)
	fmt.Printf("Slice a after change: %+v \n", aaFruits)
	fmt.Printf("Slice b after change: %v \n", bbFruits)

	fmt.Println(strings.Repeat("=", 55))

	// built in function
	fmt.Println("Values in fruits: ", len(fruits))            // take values in slice but no its memory
	fmt.Println("Total max in slice fruits: ", cap(fruits))   // take memory in slice
	fmt.Println("Total max in slice bFruits: ", cap(bFruits)) // if we do slicing with specif index ~Not 0,then the max will be not be the same with parent slice

	fmt.Println(strings.Repeat("=", 55))
	var cFruits = append(fruits, "papaya") // put the new values in the end of slice
	fmt.Printf("Add values in slice(fruits): \nFruits slice --> %v \ncFruits slice --> %v\n", fruits, cFruits)

	fmt.Println(strings.Repeat("=", 55))

}
