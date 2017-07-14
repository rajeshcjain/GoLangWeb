package main

import (
	"net"
	"fmt"
)

func main() {
	listen,_ := net.Listen("tcp",":5555")

	defer listen.Close()
	for{
		conn,_ :=listen.Accept()
		fmt.Fprint(conn," I see you connected.")
		defer conn.Close()
	}
}
