package main

import (
	"net/http"
)

// Parameter ke-1 merupakan objek untuk keperluan http response.
// Sedang parameter ke-2 yang bertipe *request ini, berisikan informasi-informasi yang berhubungan dengan http request untuk rute yang bersangkutan.

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message)) // argumen method ini ditulis dalam bentuk []type
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world"
	w.Write([]byte(message))
	// output dari setiap rute dituliskan menggunakan method write milik objek ReponseWriter
}

func main() {
	// http.HandleFunc("/", handlerIndex)      // rute / dengan aksi adalah fungsi handlerIndex()
	// http.HandleFunc("/index", handlerIndex) // rute /index dengan aksi adalah sama dengan /, yaitu fungsi handlerIndex()
	// http.HandleFunc("/hello", handlerHello) // rute /hello dengan aksi fungsi handlerHello()

	// var address = "localhost:9000"
	// fmt.Printf("served started at %s\n", address)
	// err := http.ListenAndServe(address, nil) // parameter kedua adalah objek mux atau multiplexer (kalau nill berarti pakai default)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// contoh ini, jika rute tidak didaftarkan maka akan dipanggil rute /
}
