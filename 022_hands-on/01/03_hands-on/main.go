package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/me", m)
	http.HandleFunc("/", h)

	http.ListenAndServe(":8080", nil)
}

func d(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", "Sparky")
}

func m(w http.ResponseWriter, _ *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", "Jacob")
}

func h(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Just a home page with text and no actual html or links. How boring, no?")
}
