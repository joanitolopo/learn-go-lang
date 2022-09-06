package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

func main() {
	// “Structured data” refers to data where you know the format beforehand
	birdJson := `{"species": "Pigeon", "description":"Likes to perch on rocks"}`

	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s \n", bird.Species, bird.Description)

	// Species: Pigeon, Description: Likes to perch on rocks
}
