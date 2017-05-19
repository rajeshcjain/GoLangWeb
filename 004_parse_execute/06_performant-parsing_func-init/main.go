package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

func init(){
	/*Must function is for Error checking...Here we are passing the template.ParseGlob()
	and it returns the pointer to templates which we are passing the templates
	*/
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main(){

	err := tpl.ExecuteTemplate(os.Stdout,"onegamo",nil)
	if err != nil{
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"three.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}


	err = tpl.ExecuteTemplate(os.Stdout,"two.gamo",nil)

	if err != nil{
		log.Fatalln(err)
	}


}
