package main

import (
	"net/http"
	"fmt"
)

type hotdog int

type hotcat int


func(d hotdog)ServeHTTP(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is DOG</h1>")
}

func(d hotcat)ServeHTTP(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is CAT</h1>")
}

/*
Handle registers the handler for the given pattern in the DefaultServeMux.
The documentation for ServeMux explains how patterns are matched.
*/

func main(){
	var d hotdog
	var c hotcat

	//This will give a response to localhost:8888/dog/something/anything
	//1). First way to define the handler.Here we are registering d with the default handler.So when we don't specify
	// the handler then it will be calling d and c.
	//But the problem with this we still have to declare the type and then
	//use it.So there is a better way of doing it which we have explained in the
	//next tutorial.
	http.Handle("/dog/",d)
	//This will cater on localhost:/cat all the other will not be taken care here.
	http.Handle("/cat",c)

	//Here we are specifying the nil which means that it will be
	//handled through default handler.So here the request will be going to
	//http.Handle() function where we are mentioning the Handle above.
	http.ListenAndServe(":8888",nil)
}
