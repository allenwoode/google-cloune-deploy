package main

import (
	"net/http"
	"io"
	"os/exec"
	"log"
)

func lanch() {
	cmd := exec.Command("sh", "./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is deploy server page</h1>")
	lanch()
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
	log.Println("server listening on 5000...")
}