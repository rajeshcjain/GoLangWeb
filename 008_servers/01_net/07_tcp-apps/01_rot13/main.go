package main

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func main(){

	li,err := net.Listen("tcp",":8888")

	if err != nil{
		log.Println(err)
	}

	defer li.Close()

	for{
		conn,err := li.Accept()

		if err != nil{
			log.Println(err)
		}

		go hanlde(conn)
	}
}

func hanlde(conn net.Conn){

	scan := bufio.NewScanner(conn)

	var r string
	for scan.Scan(){
		str := strings.ToLower(scan.Text())
		fmt.Println(str)
		bt := []byte(str)
		r = rot13(bt)
		fmt.Println("After the rot13 ",r)

	}
}

/*
rot13 is a format in which we send 13 character ahead of the current character at an
index.So basically encoding scheme.
*/
func rot13(bt []byte) string {

	stream := make([]byte,len(bt))

	for i,val := range bt{
		if val <= 109{
			stream[i] = val + 13
		}else{
			stream[i] = val - 13
		}
	}
	return string(stream)
}
