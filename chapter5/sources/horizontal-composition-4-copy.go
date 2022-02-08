package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func auth(s string) error {
	if s != "123" {
		return fmt.Errorf("bad auth: %s", s)
	}
	return nil
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome! %s", r.RemoteAddr)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := r.Header.Get("auth")
		if err := auth(pass); err != nil {
			log.Println("bad auth: ", pass)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%v [%s] %q from %s", time.Now(), r.Method, r.URL.Path, r.RemoteAddr)
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":7000", logHandler(authHandler(http.HandlerFunc(greeting))))
}
