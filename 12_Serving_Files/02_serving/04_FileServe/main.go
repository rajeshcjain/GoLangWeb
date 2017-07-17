package main

import (
	"net/http"
	"io"
)

func handleImageReq(w http.ResponseWriter,req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src = "/Golde33443.jpg">`)
}

func main() {
	/*
	This is a FileServer which accepts the FileSystem interface and
	 This interface accepts the open function which is implemented by
	 http.Open and when we pass it then it will open the directory.

So here both the files will be served that means main.go and the other file in the
directory.
	*/
	http.Handle("/",http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog",handleImageReq)
	http.ListenAndServe(":7777",nil)
}
