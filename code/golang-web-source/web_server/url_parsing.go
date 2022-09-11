package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://dasarpemrogramangolang.novalagung.com/A-url-parsing.html"
	var u, e = url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)

}
