package main

import (
	"text/template"
	"log"
	"os"
)

func main(){
	tpl,err := template.ParseGlob("template/*")

	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"two.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"three.html",nil)

	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"one.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}

}
