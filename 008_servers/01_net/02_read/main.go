package main

import (
	"net"
	"io/ioutil"
	"fmt"
	"bufio"
)

func main(){

	li,err := net.Listen("tcp",":8888")

	if err != nil{
		panic(err)
	}
	defer li.Close()

	for{
		conn,err := li.Accept()
		if err != nil{
			panic(err)
		}
		go handleInputConn(conn)

	}


}

func handleInputConn(conn net.Conn){
	scan := bufio.NewScanner(conn)
	for scan.Scan(){
		ln := scan.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	//This line will not be executed as scanner will
	//not return.scan() function return with false when it reaches to
	//end of line.But in this case it will keep on listening
	//input stream for ever.Like keep on asking are you still
	//there as we do on cell phone when we don't listen anything
	//from other side.
	//But this is not the case with other function as they read the data
	//and then return with the result as in case with ReadAll() function

	fmt.Println("This line will never be executed.")

}


/*
The problem with the below code is that it will just read just once
and then it will be quite.where as scanner will keep on listening
and asking are you still there.

*/
func handleInputConnUsingReadAll(conn net.Conn ){
	stream,err := ioutil.ReadAll(conn)

	if err != nil{
		panic(err)
	}

	fmt.Println(string(stream))
	defer conn.Close()
}
