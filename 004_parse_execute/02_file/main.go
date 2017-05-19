package main

import (
	"text/template"
	"log"
	"os"
)

func main(){

	tpl,err := template.ParseFiles("go.html")

	if err != nil{
		log.Fatalln(err)
	}

	//Here we are creating the file called index.
	fDescriptor,err := os.Create("index.html")

	if err != nil{
		log.Fatalln(err)
	}

	//And here we are directing the content to index file....through the pointer which
	//points to the location where the template is kept.
	tpl.Execute(fDescriptor,nil)
}
