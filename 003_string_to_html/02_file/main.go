package main

import (
	"fmt"
	"os"
	"io"
	"strings"
)

func main(){

	name := "Rajesh Jain"
	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(template)
	str := fmt.Sprint(template)

	fmt.Println(str)

	//This command os used for creating the file.So here we are creating the
	//index.html which can be used for I/O.The first return value is the file descriptor and
	//2nd value is error.
	fDescriptor, _ := os.Create("index.html")

	//Deferring it till the main function termination...till then we can use
	//fDescriptor easily.
	defer fDescriptor.Close()

	//So here we learning something really good..we are seeing packages like
	//os and below package for io and strings.Need to learn more these
	//packages
	io.Copy(fDescriptor,strings.NewReader(str))
}
