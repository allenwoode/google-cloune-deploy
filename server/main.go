package main

import (
	"io"
	"net/http"
	"log"
)

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>server v4.0 index page!</h1>")
	log.Println("request index page")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}