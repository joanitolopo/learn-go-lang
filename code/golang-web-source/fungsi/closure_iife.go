package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//  Provide seed
	rand.Seed(time.Now().Unix())

	numbers := rand.Perm(10)

	newNumbers := func(min int) (r []int) {
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return
	}(3)

	fmt.Println("Originial number: ", numbers)
	fmt.Println("Filtered number: ", newNumbers)

}
