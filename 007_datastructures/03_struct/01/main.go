package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("go.html"))
}


//This is how you create a struct...that means its a new type
type sage struct {
	Name  string
	Motto string
}

func main(){

	//Now here we are assigning new value to that struct and
	//giving it to buddha
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout,buddha)

	if err != nil{
		log.Fatalln(err)
	}



}
