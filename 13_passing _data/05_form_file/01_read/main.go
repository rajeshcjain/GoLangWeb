package main

import (
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/",handleReq)
	http.ListenAndServe(":7777",nil)
}

func handleReq(w http.ResponseWriter,req *http.Request){

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost{
		f,h,err := req.FormFile("q")

		if err != nil {
			http.NotFound(w,req)
		}
		defer f.Close()

		fmt.Println("file \t",f,"\t header \t",h,"\t err \t",err)
		str,_ := ioutil.ReadAll(f)
		s = string(str)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
