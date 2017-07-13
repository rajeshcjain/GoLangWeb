package main

import (
	"net"
	"io/ioutil"
	"fmt"
)

func main(){
	/*
	The difference between Listen and Dial function is that
	Dial can be used for establishing the connection with the
	server(which means where we have created the server.)
	So when we Dial in to some server.It creates the connection
	with that server and then operate
	*/

	/*
	To see it working we need to run the write server which we
	implemented in 01.So run both the programs in ||ly
	*/

	conn,err := net.Dial("tcp",":8888")

	if err != nil{
		panic(err)
	}
	defer conn.Close()

	data,err := ioutil.ReadAll(conn)
	fmt.Println(string(data))
}