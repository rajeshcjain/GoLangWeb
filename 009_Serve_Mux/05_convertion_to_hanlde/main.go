package main

import (
	"net/http"
	"fmt"
)

type hotdog int

type hotcat int


func d(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is DOG!!</h1>")
}

func c(w http.ResponseWriter,req * http.Request){
	fmt.Fprint(w,"<h1>The response is CAT!!</h1>")
}

/*
Handle registers the handler for the given pattern in the DefaultServeMux.
The documentation for ServeMux explains how patterns are matched.
*/

func main(){

	//This is a 3rd way of doing it.
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//This is a important concept here we have to understand...here we are converting
	//the function func(w http.ResponseWriter,req * http.Request) to
	//http.HandlerFunc(c)..there is a type which is declared in http package
	//which is of type HandlerFunc


	/*
	type HandlerFunc func(ResponseWriter, *Request)
	and we have associated function with it as

	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
	}

so as this has ServeHTTP() function implemented this function is implementing the
Handle interface...so this will work
	*/

	/*
To understand this concept more lets consider the below example

package main

import (
	"fmt"
)

type hotdog int

func main() {
	var x int
	x = 7

	var y hotdog
	y = hotdog(x)

	fmt.Println(y)
}
The above code will work as we can convert the type hotdog to int.In this as the underlying
types are same so the conversion can happen.The same thing is happening in
HanldeFunc which has a under lying type as func(w ResponseWriter,req *Request);So
the conversion can happen.


But the below code will not work as hotdog and int are of different type.

package main

import (
	"fmt"
)

type hotdog int

func main() {
	var x hotdog
	x = 7

	var y int
	y = int(x)

	fmt.Println(y)
}




	*/


	//This will give a response to localhost:8888/dog/something/anything
	http.Handle("/dog/",http.HandlerFunc(d))

	//This will cater on localhost:/cat all the other will not be taken care here.
	http.Handle("/cat",http.HandlerFunc(c))

	//Here we are specifying the nil which means that it will be
	//handled through default handler.So here the request will be going to
	//http.Handle() function where we are mentioning the Handle above.
	http.ListenAndServe(":8888",nil)
}
