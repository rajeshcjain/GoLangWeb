package main

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func main() {

	li, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Println(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

var myMap = make(map[string]string)

func handle(conn net.Conn){

	scan := bufio.NewScanner(conn)
	for scan.Scan(){

		input := scan.Text()
		fmt.Println(input)
		str := strings.Fields(input)
		simulateRedis(conn,str)
	}
}

func simulateRedis(conn net.Conn,str []string) {

	fmt.Println("In simulateRedis",str)

	switch str[0] {
	case "Get":
		fmt.Fprintf(conn,myMap[str[1]])
	case "Set":
		if len(str) != 3{
			fmt.Fprintf(conn,"There is a problem with input")
		}
		myMap[str[1]] = str[2]
	case "Del":
		delete(myMap,str[1])
	default:
		fmt.Fprintf(conn,"invalid command")
	}
}