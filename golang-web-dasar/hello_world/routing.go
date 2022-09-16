package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Penggunaan http.HandleFunc()
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello again"))
	})

	fmt.Println("server starter at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
