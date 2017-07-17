package main

import (
	"net/http"
	"os"
)


func hanldeImageReq(w http.ResponseWriter, Req *http.Request) {
	f, err := os.Open("Golde33443.jpg")
	if err != nil {
		http.Error(w, "Not Found", 404)
		return
	}

	fi, _ := f.Stat()

	http.ServeContent(w, Req, fi.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/img",hanldeImageReq)
	http.ListenAndServe(":7777",nil)
}
