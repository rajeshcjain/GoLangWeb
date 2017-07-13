package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

/*
                        IMPORTANT

   This is important,Here in "tpl.html"..we are representing the data by wisdom variable.

*/
func init(){
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}