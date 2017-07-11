package main

import (
//	"github.com/alecthomas/template"
	"os"
	"log"
	"text/template"
)

/* Go in godoc.com and search for "template".Now select "text/template".
Go To index and see "type Template".We have functions like Must,New,ParseFiles,ParseGlob which
returns pointer to Template.

There are function which are like

func (t *Template) AddParseTree(name string, tree *parse.Tree) (*Template, error)

where  (t *Template) is a pointer to template and you can think of it as a pointer to a template container
where different templates are parsed and held.

*/
func main(){

	//Here we are learning of template package...we are getting the parseFiles...
	//It parses the input file and and gives the Template.
       tpl,err := template.ParseFiles("go1.html","go.html")

	if err != nil {
		// This is important...Here we are learning new thing.It is equivalent to
		// fmt.println() + os.Exit(1)
		log.Fatalln(err)
	}

	// Execute applies a parsed template to the specified data object,
	// and writes the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.

	// Here we are not providing the data as we don't want to change any data in template..
	//it will print the content on the Stdout.
	err = tpl.Execute(os.Stdout,nil)
	if err != nil {
		log.Fatalln(err)
	}

}
