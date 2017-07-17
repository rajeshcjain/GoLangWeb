package main

import (
	"net/http"
	"io"
)


func handleImageReq(w http.ResponseWriter,req * http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="/Golde33443.jpg">`)
}

func main() {

	http.HandleFunc("/img",handleImageReq)
	http.ListenAndServe(":7777",nil)

}
