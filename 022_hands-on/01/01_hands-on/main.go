package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", m)

	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Woof woof! I own you!")
}

func m(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Welcome back, J-Goody Towns!")
}
