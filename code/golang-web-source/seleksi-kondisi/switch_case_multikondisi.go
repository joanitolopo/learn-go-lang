package main

import "fmt"

func main() {
	var point = 2

	switch point {
	case 8:
		fmt.Println("Perfect")
	case 7, 6, 5, 4:
		fmt.Println("Awesome")

	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You can be better!")
		}

	}
}
