package main

import "fmt"

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d divider by %d\n", m, n)
		return // if we dont return the value then go will exit each process after that
	}

	res := m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}
