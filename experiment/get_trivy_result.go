package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Vulns struct {
	Vulns []Vuln `json:"Results"`
}

type Vuln struct {
	Target string `json:"Target"`
	Class  string `json:"Class"`
	Type   string `json:"Type"`
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("data/result.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened results.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var vulns Vulns

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &vulns)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(vulns.Vulns); i++ {
		fmt.Println("Target: " + vulns.Vulns[i].Target)
		fmt.Println("Class: " + vulns.Vulns[i].Class)
		fmt.Println("Type: " + vulns.Vulns[i].Type)
		fmt.Println()
	}

}
