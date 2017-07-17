package main

import "net/http"


/*
In this program,we dont want any other file to be served
*/



func main() {

	/*
	This is to show that when we call the the request with / then
	 it will return it from the main directory that means it will
	 show main.go and assets.
	 whereas when called with
	  http.Handle("/resource/",http.FileServer(http.Dir(".")))

	  then it will not serve any thing.
	  Please see the FileServer documentation.
	  */

	//This request will work
	http.Handle("/",http.FileServer(http.Dir("./assets")))

	/*
	But if we want to make it little meaningfull then we can strip
	the incoming request and use it as above.

	http.Dir("./assets") meaning
	"." --> from the home directory start.

	*/
        http.Handle("/resource/",http.StripPrefix("/resource",http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":7777",nil)
}
