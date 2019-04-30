package main

import (
	"io"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/assets/dog.jpg", doggy)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tmpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func doggy(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "assets/dog.jpg")
}
