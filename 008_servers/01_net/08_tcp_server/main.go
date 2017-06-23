package main

import (
	"net"
	"bufio"
	"strings"
	"fmt"
)

func main(){

	li,err :=net.Listen("tcp",":8888")

	if err != nil{
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil{
			panic(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn){

	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn){

	scan := bufio.NewScanner(conn)

	firstLine := 0

	for scan.Scan(){
		str := scan.Text()
		fmt.Println(str)
		if firstLine == 0{
			firstStr := strings.Fields(str)[0]
			secStr := strings.Fields(str)[1]
			fmt.Printf("****Method %s\n",firstStr)
			fmt.Printf("****URI %s\n",secStr)
		}
		if str == ""{
			break
		}
		firstLine++
	}
}

func response(conn net.Conn){
	res := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(res))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprintf(conn,res)

}


