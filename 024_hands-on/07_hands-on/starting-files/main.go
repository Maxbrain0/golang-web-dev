package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fs := http.StripPrefix("/resources/", http.FileServer(http.Dir("public")))
	http.HandleFunc("/", home)
	http.Handle("/resources/pics/", fs)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
