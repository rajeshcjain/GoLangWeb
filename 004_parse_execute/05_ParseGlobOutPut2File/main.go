package main

import (
	"os"
	"log"
	"text/template"
)

func main(){

	tpl,err := template.ParseGlob("template/*")

	if err != nil{
		log.Fatalln(err)
	}

	fDesc,_ := os.Create("index.output")

	err = tpl.ExecuteTemplate(fDesc,"two.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(fDesc,"three.html",nil)

	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(fDesc,"one.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}


}
