package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("go.html"))
}


func main(){
	sages := map[string]string{
		"rajesh" : "jain",
		"sakshi" : "jain",
		"ramesh" : "jain",
		"Atishay" : "Jain",
	}

	err := tpl.Execute(os.Stdout,sages)

	if err != nil{
		log.Fatalln(err)
	}

}
