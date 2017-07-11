package main

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("I am in ServeHTTP")

	/*
	ParseForm is the method which needs to be call before we can use
	req.Form  - This for getting the data in the body
	req.PostForm - this is for getting the data from body and header
	*/

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(req.Form)

	tpl.ExecuteTemplate(w, "index.html", req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d hotdog
	http.ListenAndServe(":7777", d)
}