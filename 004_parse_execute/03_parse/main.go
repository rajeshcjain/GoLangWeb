package main

import (
	"text/template"
	"log"
	"os"
)

func main(){

	/*Here we are parsing the file and putting it in to template and it returns the
	  template pointer.Now when we call tpl.Execute() then this will call the first
	  template which is available in tpl pointer as we know we can think of tpl as
	  the pointer to space where template are kept.
	*/
	tpl,err := template.ParseFiles("one.gmao")
	if err != nil{
		log.Fatalln(err)
	}
	err = tpl.Execute(os.Stdout,nil)

	if err != nil{
		log.Fatalln(err)
	}

	/*
	Now we want to keep add the template in to tpl..so we call the
         func (t *Template) ParseFiles(filenames ...string) (*Template, error)
         which is a function which is being called on the pointer which is returned from
         the pointer where we keep the templates.

	*/
	tpl,err = tpl.ParseFiles("two.gmao","three.gmao")
	if err != nil {
		log.Fatalln(err)
	}

	/*
	  If we have more then onw template then we can call it through
	  func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{})
	  where we specify the template we want to print on the output.
	*/
	err = tpl.ExecuteTemplate(os.Stdout,"two.gmao",nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"three.gmao",nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gmao", nil)
	if err != nil {
		log.Fatalln(err)
	}

	/*
	 Now we have multiple templates and when we call the Execute function in that case it will print the
	 first template which is available...and in this case it is one.gmao
	*/
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
