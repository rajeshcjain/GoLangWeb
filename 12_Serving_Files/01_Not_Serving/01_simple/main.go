package main

import (
	"net/http"
	//"fmt"
	"io"
)


func handleReq(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//fmt.Fprint(w,`<img>https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg</img>`)
	io.WriteString(w, `
	<!--not serving from our server-->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
	`)
}

func main() {
	http.HandleFunc("/img",handleReq)
	http.ListenAndServe(":7777",nil)
}
