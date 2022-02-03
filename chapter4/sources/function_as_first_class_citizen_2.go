package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang\n")
	fmt.Fprintf(w, "Hello Golang, you are from %s\n", r.URL.Path)
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
	// http.ListenAndServe(":8080", greeting)
}
