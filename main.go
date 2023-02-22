package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", myMiddleware(myHandler))
	http.ListenAndServe(":8082", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello from Handler")
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello from Middleware")
		next(w, r)
	}
}
