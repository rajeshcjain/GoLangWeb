package main

import (
	"net"
	"bufio"
	"time"
	"fmt"
)

func main(){
        li,err := net.Listen("tcp",":8888")

	if err != nil{
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil{
			panic(err)
		}

		go hanlde(conn)
	}
}

func hanlde(conn net.Conn){

	//Here we are setting up the deadline time for the request

	err := conn.SetDeadline(time.Now().Add(time.Second*10))

	if err != nil{
		fmt.Println("Connection Timed Out")
	}

	scan := bufio.NewScanner(conn)
	for scan.Scan(){
		str := scan.Text()
		fmt.Println(str)
		fmt.Fprint(conn,"The content i got is ",str)
	}

	fmt.Println("So At the end of the Program")
}
