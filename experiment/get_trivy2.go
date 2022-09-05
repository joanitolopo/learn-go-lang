package main

import (
	"fmt"

	"github.com/Jeffail/gabs"
)

func main() {

	jsonFile, _ := gabs.ParseJSONFile("data/result.json")
	value := jsonFile.Path("Results.Target")

	fmt.Println(value)

}
