package main

import (
	"log"
	"net/http"
)

func main() {

	// public views

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/user/create", createUser(HandleCreateUser))
	http.HandleFunc("/user/login", loginUser(HandleLoginUser))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
