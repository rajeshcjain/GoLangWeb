package main

import (
	"os"
	"fmt"
	"log"
	"io"
	"strings"
)

func main(){
	// So here we are learning something very important.
	//In the os package...how can we provide command line arguments...
	// os provides Args slice which is used for getting the number of arguments
	//passed which are separated by space.So os.Args[0] is program name and os.Args[1] is
	// parameter passed after the program name.

	// so here we will run it by "go run main.go RajeshJain
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` +
		name +
		`</h1>
		</body>
		</html>
	`)
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
