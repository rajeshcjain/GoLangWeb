package main

import "net"

func main(){
	listener,_ := net.Listen("tcp",":8080")

	defer listener.Close()

	for {
		conn,_ := listener.Accept()
		conn.Write([]byte("Hello world!!"))
		conn.Close()
	}




}
