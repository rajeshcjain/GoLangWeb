1). Take the previous program and change it so that:

a). func main uses http.Handle instead of http.HandleFunc
         Contstraint: Do not change anything outside of func main

Hints:

http.HandlerFunc
           type HandlerFunc func(ResponseWriter, *Request)


http.HandleFunc
          func HandleFunc(pattern string, handler func(ResponseWriter, *Request))


source code for HandleFunc

  func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
  		mux.Handle(pattern, HandlerFunc(handler))
  }