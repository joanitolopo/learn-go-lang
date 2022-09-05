package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	max := 3
	numbers := rand.Perm(11)
	howMany, getNumbers := findMax(numbers, max)
	theNumbers := getNumbers()

	fmt.Println("Numbers\t:", numbers)
	fmt.Printf("Find \t: %d\n\n", max)
	fmt.Println("Found \t:", howMany)
	fmt.Println("Value \t:", theNumbers)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int

	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}

}
