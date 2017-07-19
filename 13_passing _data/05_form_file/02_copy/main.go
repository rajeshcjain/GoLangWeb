package main

import (
	"text/template"
	"net/http"
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("template/index.html"))
}

func main() {

	http.HandleFunc("/",handleReq)
	http.Handle("/favicon",http.NotFoundHandler())
	http.ListenAndServe(":7777",nil)
}

func handleReq(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Method)
	var s string

	if req.Method == http.MethodPost{
		f,h,err := req.FormFile("q")

		if err != nil{
			http.NotFound(w,req)
		}

		fmt.Println("file \t ",f, "\t header \t ",h,"\t err \t",err)
		defer f.Close()

		fd,_:= os.Create(filepath.Join("./user/",h.Filename))

		defer fd.Close()

		byteData,_ := ioutil.ReadAll(f)
		s = string(byteData)

	}

	tpl.ExecuteTemplate(w,"index.html",s)

}
