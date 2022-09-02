package main

import "fmt"

func main() {
	var point = 3

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8) && (point > 3):
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You need to learn more.")
		}
	}
}
