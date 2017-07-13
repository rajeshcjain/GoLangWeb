package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter,req *http.Request){
	switch req.URL.Path{
	case "/dog" :
		fmt.Fprint(w,"<h1>DOG</h1>")
	case "/cat" :
		fmt.Fprint(w,"<h1>CAT</h1>")
	}
}

func main(){
	var d hotdog
	http.ListenAndServe(":8080",d)
}
