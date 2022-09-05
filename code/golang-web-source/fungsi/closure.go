package main

import "fmt"

func main() {
	// Closure adalah sebuah fungsi yang bisa disimpan dalam variabel
	getMinMax := func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return
	}

	numbers := []int{2, 3, 4, 3, 4, 2, 3}
	min, max := getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers, min, max)

	// ouput:
	// data : [2 3 4 3 4 2 3]
	// min : 2
	// max : 4

}
