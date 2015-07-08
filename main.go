package main

import (
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func main() {

	// public views

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/user/create", createUser(HandleCreateUser))
	http.HandleFunc("/user/login", loginUser(HandleLoginUser))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
