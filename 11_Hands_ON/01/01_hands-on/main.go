package main

import (
	"net/http"
	"fmt"
)


func handleDefault(w http.ResponseWriter,Req *http.Request){
	fmt.Fprint(w,"<h1>Handle the Default</h1>")
}

func hanldeDogfunc(w http.ResponseWriter,Req *http.Request){
	fmt.Fprint(w,"<h1>Handle the Dog Request</h1>")
}

func hanldeNameReq(w http.ResponseWriter,Req *http.Request){
	fmt.Fprint(w,"<h1>Handle the name Request,name is Raj</h1>")
}

func main() {


	http.HandleFunc("/",handleDefault)
	http.HandleFunc("/dog/",hanldeDogfunc)
	http.HandleFunc("/me/",hanldeNameReq)
	http.ListenAndServe(":8080",nil)
}
