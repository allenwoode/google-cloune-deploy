package main

import (
	"net/http"
	"fmt"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, this is index page</h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
	fmt.Println("server listening on 8000...")
}