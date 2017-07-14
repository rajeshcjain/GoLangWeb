package main

import (
	"text/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("sample.html"))
}

func main() {
	http.Handle("/",http.HandlerFunc(handleDefault))
	http.Handle("/dog/",http.HandlerFunc(handleDogFunc))
	http.Handle("/me/",http.HandlerFunc(handleNameReq))
	http.ListenAndServe(":7777",nil)
}

func handleDefault(w http.ResponseWriter,Req *http.Request){
	fmt.Fprint(w,"<h1>Handle the Default</h1>")
}

func handleDogFunc(w http.ResponseWriter,Req *http.Request){

	fmt.Fprint(w,"<h1>Handle the Dog Request</h1>")
}

func handleNameReq(w http.ResponseWriter,Req *http.Request){
	tpl.ExecuteTemplate(w,"sample.html","Raj")
}
