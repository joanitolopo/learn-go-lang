package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var baseURL = "http://localhost:8080"

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func users(w http.ResponseWriter, r *http.Request) { //fungsi untuk handle endpoint /users (mengembaikan semua sample data dalam array)
	w.Header().Set("Content-Type", "application/json") // menentukan tipe response yaitu sebagai json

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result) // mendaftarkan data sebagai response
		return
	}
}

func user(w http.ResponseWriter, r *http.Request) { // endpoint untuk mengembalikan satu buah data saja
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id") // mengambil data form yang dikirim
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}
		}

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhot:8080/")
	http.ListenAndServe(":8080", nil)
}
