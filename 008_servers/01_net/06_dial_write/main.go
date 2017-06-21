package main

import (
	"net"
	"io"
	"fmt"
)

func main(){
	/*
	Here is the Dial a connection with the read server..which we wrote
	in 02_read program.so run both the programs in ||ly to see it working
	*/
	conn,err := net.Dial("tcp",":8888")

	if err != nil{
		 panic(err)
	}

	defer conn.Close()

	io.WriteString(conn,"I am doing great you tell me.")
	fmt.Fprintf(conn,"How is life treating you.")
}
