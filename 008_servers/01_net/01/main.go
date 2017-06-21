package main

import (
	"net"
	"io"
	"fmt"
)

func main(){

	conn,err := net.Listen("tcp",":8888")
	if err != nil{
		panic(err)
	}

	defer conn.Close()

	for {
		con, err := conn.Accept()

		if err != nil {
			panic(err)
		}
		//This method send the data to stream
		io.WriteString(con,"\n Hello From TCP Server")
		fmt.Fprint(con,"How is your Day?")
		con.Close()
	}

}
