package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Wick"}
	printMessage("Halo", names) // Halo John Wick

}

func printMessage(message string, arr []string) {
	fmt.Println("Slice Value: ", arr)
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)

}
