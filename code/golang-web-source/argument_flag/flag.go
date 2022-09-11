package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type ypur age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)

	// cara ke dua
	var data2 string
	flag.StringVar(&data2, "gender", "male", "type your gender")

	// contoh run:
	// run flag.go -name="john wick" -age=28

	// source: https://dasarpemrogramangolang.novalagung.com/A-command-line-args-flag.html
}
