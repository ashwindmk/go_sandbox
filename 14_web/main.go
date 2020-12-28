package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about me")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server started...")
	http.ListenAndServe(":3000", nil)
}
