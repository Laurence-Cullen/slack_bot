package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bot", botHandler) // each request calls staticHandler
	//http.HandleFunc("/register_user", registerUser)

	log.Fatal(http.ListenAndServe(":8500", nil))
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // Parses the request body
	challenge := r.PostForm.Get("challenge")
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprint(w, challenge)
}