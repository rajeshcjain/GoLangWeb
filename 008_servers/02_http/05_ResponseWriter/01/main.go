package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (d hotdog) ServeHTTP(w http.ResponseWriter,req *http.Request){

	w.Header().Set("test1","value1")
	w.Header().Set("test2","value2")
	fmt.Fprint(w,"<h1>This is the output</h1>")
}

func main(){
	var d hotdog
	http.ListenAndServe(":8888",d)
}
