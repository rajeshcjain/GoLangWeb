package main

import (
	"io"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":7777", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	/*
	So lets understand this..here when we press localhost:7777/ on the
	browser then the request comes which does not have "q" as a query param
	so
	v := req.FormValue("q")
	will give null. so it will print

	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)

	which will be a text box and as v is null so nothing will be printed
	now when we enter the value in text box and the value is given name
	"q".The request comes again to this function and it will try to get the
	value of "q" using
	v := req.FormValue("q")
	this time it will get the value and it will get it in body of the
	request as it is is a post request when user press the submit button
	of the request.
	now with the text box the value is also printed.

	*/
	v := req.FormValue("q")
	fmt.Println("The requested value is ",v)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
}
