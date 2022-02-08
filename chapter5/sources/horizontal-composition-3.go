package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, fmt.Sprintf("Hi, you access: %s", r.URL))
	// err = errors.New("generate an error")
	if err != nil {
		log.Println("err: ", err)
	}
}

func main() {
	http.ListenAndServe(":7000", http.HandlerFunc(greeting))
}
