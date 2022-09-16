package main

import (
	"fmt"
	"net/http"
)

func main() {
	// penggunaan http.Handle() untuk routing static assets
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assests"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
