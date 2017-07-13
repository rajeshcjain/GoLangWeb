package main

import (
	"github.com/julienschmidt/httprouter"
	"text/template"
	"net/http"
)


var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main(){

	mux := httprouter.New()
	mux.GET("/",index)
	mux.GET("/about",about)
	mux.GET("/contact",contact)
	mux.GET("/apply",apply)
	mux.POST("/apply",applyProcess)
	mux.GET("/user/:name",user)
	mux.GET("/user/:category/:article",blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", mux)
}

func blogWrite(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}

func blogRead(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}


func user(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}


func applyProcess(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}


func apply(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}


func contact(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}

func about(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}

func index(w http.ResponseWriter,req *http.Request,ps httprouter.Params){

}
