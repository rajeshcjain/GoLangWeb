package main

import (
	"net/http"
	"io"
)

func handleQuery(w http.ResponseWriter,req *http.Request){
	// FormValue returns the first value for the named component
	// of the query.POST and PUT body parameters take precedence
	// over URL query string values.
	//It gives the value when queried from
	//
	// localhost:8080/?q=dog
	//
	// now the value of v will be dog
	v := req.FormValue("q")
        io.WriteString(w,"The search is " +v)
}

func main() {

	http.HandleFunc("/",handleQuery)
	http.ListenAndServe(":7777",nil)
}
