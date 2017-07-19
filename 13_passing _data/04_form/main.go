package main

import (
	"text/template"
	"net/http"
)


//Define the data structure to get the data from the form entered.
type person struct {
   FirstName string
   LastName string
    Subscribed bool
}

 var tpl *template.Template

func init(){
	//parse the templates
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func handleReq(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	/*
	When the user hit the submit button then user can get the value
	otherwise it will be null and in case of subscribe it will be false
	as it will not be equal to "on".
	*/
	firstName := req.FormValue("first")
	lastName := req.FormValue("last")
	sub := req.FormValue("subscribe") == "on"

        per := person{firstName,lastName,sub}

	tpl.ExecuteTemplate(w,"index.html",per)

}

func main(){
	http.HandleFunc("/",handleReq)
	http.ListenAndServe(":7777",nil)
}
