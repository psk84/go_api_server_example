package main

import (
	"io"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world\n")
}
