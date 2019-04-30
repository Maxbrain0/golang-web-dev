package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("starting-files"))

	http.Handle("/", fs)

	log.Println("Listining on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
