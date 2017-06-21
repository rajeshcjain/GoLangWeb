package main

import (
	"net"
	"bufio"
	"fmt"
)

func main(){

	li,err := net.Listen("tcp",":8888")

	if err != nil{
		panic(err)
	}

	for{
		conn,err := li.Accept()
		if err != nil{
			panic(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn){
	scan := bufio.NewScanner(conn)

	for scan.Scan(){
		ln := scan.Text()
		fmt.Println(ln)
	/*Now here is a important information.When we send the response
	back to browser.Browser request get full filled and it
	closes the connection and in that case scan() function
	will return with false and hence the connection will be closed
	*/
		/*
		So better test with telnet.
		Start the program and the run the telnet on separate
		command prompt and then
		telnet localhost 8888

		I have heard you : hello
		I have heard you : how are you doing
		I have heard you : how are you doing    sldjsladj;slajkf
		I have heard you :      sldjsladj;slajkf

		*/

		fmt.Fprintf(conn,"I have heard you : %s",ln)
	}
	defer conn.Close()

	//The below line will never be executed as scan will always
	//keep on listening
	fmt.Println("This will not be executed")

}
