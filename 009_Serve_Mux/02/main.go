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

func main(){
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	//This will give a response to localhost:8888/dog/something/anything
	mux.Handle("/dog/",d)
	//This will cater on localhost:/cat all the other will not be taken care here.
	mux.Handle("/cat",c)

	//Here we are giving the Handle to ListenAndServe;So we are giving the
	// our handle to the ListenAndServe which will be taking care of the
	//incoming requests.
	http.ListenAndServe(":8888",mux)
}
