package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type newType struct {
       str1 string
       str2 string
}

func init(){
	tpl = template.Must(template.ParseFiles("go.html"))
}

func main(){

	sages := newType{
		str1 : "this is a new type str1",
		str2 : "This is a new type str2",
	}

	err := tpl.Execute(os.Stdout,sages)

	if err != nil{
		log.Fatalln(err)
	}

}
