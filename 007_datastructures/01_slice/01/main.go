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

func main(){

	sages := []string {"Rajesh Jain","Sakshi Jain","Atishay Jain","Ramesh Chand Jain","Kamlesh Jain"}
	err :=tpl.Execute(os.Stdout,sages)

	if err != nil{
		log.Fatalln(err)
	}
}


