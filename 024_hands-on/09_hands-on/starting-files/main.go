package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	fs := http.FileServer(http.Dir("./public/"))
	http.HandleFunc("/", home)
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
