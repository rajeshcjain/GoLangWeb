package main

import (
	"net/http"
	"fmt"
)

type hotdog int

type hotcat int


func d(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is DOG!!</h1>")
}

func c(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is CAT!!</h1>")
}

/*
Handle registers the handler for the given pattern in the DefaultServeMux.
The documentation for ServeMux explains how patterns are matched.
*/

func main(){

	//This is a 2nd way of doing it.
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//is a function definition and found that it accepts the function not Handle.
	//So we just need to pass the function here which will handle the incoming
	//request.

	//This will give a response to localhost:8888/dog/something/anything
	http.HandleFunc("/dog/",d)
	//This will cater on localhost:/cat all the other will not be taken care here.
	http.HandleFunc("/cat",c)

	//Here we are specifying the nil which means that it will be
	//handled through default handler.So here the request will be going to
	//http.Handle() function where we are mentioning the Handle above.
	http.ListenAndServe(":8888",nil)
}
