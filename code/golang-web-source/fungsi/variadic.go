package main

import (
	"fmt"
	"strings"
)

func main() {
	// variadic function atau pembuatan fungsi dengan parameter sejenis yang tak terbatas

	avg := calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("Rata-rata: %.2f", avg)
	fmt.Println("Initialize data with int --> ", msg)

	fmt.Println(strings.Repeat("=", 50))

	numbers := []int{2, 4, 3, 5, 3, 2, 2, 5, 5, 3}
	avg2 := calculate(numbers...)
	msg2 := fmt.Sprintf("Rata-rata : %.2f", avg2)

	fmt.Println("Initialize data with slice -->", msg2)

}

func calculate(numbers ...int) (avg float64) {
	total := 0

	for _, number := range numbers {
		total += number
	}

	fmt.Println("Numbers with 0-th index: ", numbers[0])

	avg = float64(total) / float64(len(numbers))
	return

}
