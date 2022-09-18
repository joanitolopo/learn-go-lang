package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fungsi untuk render html
		var filepath = path.Join("views", "index.html")
		var tmpl, err = template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "learning golang web",
			"name":  "batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets")))) // parsing css

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
