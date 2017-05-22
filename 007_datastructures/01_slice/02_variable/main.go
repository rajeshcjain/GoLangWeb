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


/*
This is nice...we are passing the slice here and in the go.html..we are having
2 elements where one is index and other is element and we range it through
*/
func main(){

	sages := []string {"Rajesh Jain","Sakshi Jain","Atishay Jain","Ramesh Chand Jain","Kamlesh Jain"}
	err :=tpl.Execute(os.Stdout,sages)

	if err != nil{
		log.Fatalln(err)
	}
}


