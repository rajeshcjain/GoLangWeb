package main

import "net/http"


/*
This will serve the Single file.FileServer is used for serving the multiple
Files.
*/
func hanldeImageReq(w http.ResponseWriter,req *http.Request){
	http.ServeFile(w,req,"Golde33443.jpg")

}

func main() {

	http.HandleFunc("/img",hanldeImageReq)
	http.ListenAndServe(":7777",nil)
}
