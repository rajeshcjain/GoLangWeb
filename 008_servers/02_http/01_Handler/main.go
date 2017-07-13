package main

import (
	"net/http"
	"fmt"
)

type hotdog int

func (d hotdog) ServeHTTP(response http.ResponseWriter,r *http.Request){
	fmt.Fprintf(response,"HI I am here")
}


func main(){
	var d hotdog
	http.ListenAndServe(":7777",d)
}
