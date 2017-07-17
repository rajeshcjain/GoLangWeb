package main

import (
	"net/http"
	"io"
	"os"
)

func dog(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src = "/Golde33443.jpg">`)
}

/*
This is the function which will do the copy from the source to destination

*/
func dogFunc(w http.ResponseWriter,req *http.Request){
	//Open the file.
	fDiscriptor,err := os.Open("Golde33443.jpg")

	if err != nil{
		http.Error(w,"File not Found.",404)
	}

	defer fDiscriptor.Close()
	//Here we copy the file from Destination to source.
	io.Copy(w,fDiscriptor)
}

func main() {
	http.HandleFunc("/",dog)
	http.HandleFunc("/img",dogFunc)
	http.ListenAndServe(":7777",nil)
}
