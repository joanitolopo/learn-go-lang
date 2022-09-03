package main

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

// registered http request handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// used to create an http request multiplexer
	// request multiplexer matches the url of incoming request against a list of registered patterns
	// and then calls the associated handler for the pattern whenever match is found
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":"+port, mux)

}
