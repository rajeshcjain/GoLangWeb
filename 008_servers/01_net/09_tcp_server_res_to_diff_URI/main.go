package main

import (
	"net"
	"bufio"
	"fmt"
	"strings"
	//"sync"
)

//var wg sync.WaitGroup

func main(){


	li,err := net.Listen("tcp",":7777")

	if err != nil{
		panic(err)
	}
	defer li.Close()

	for{
		conn,err := li.Accept()
		if err != nil{
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn){
	//wg.Add(1)
	defer conn.Close()
	request(conn)
	//wg.Wait()
}

func request(conn net.Conn){

	scan := bufio.NewScanner(conn)
	firstLine :=0
	var listOfStrings []string
	var m string
	var uri string
	for scan.Scan(){
		str := scan.Text()
		fmt.Println(str)
		if firstLine == 0 {
			listOfStrings = strings.Fields(str)
			m = listOfStrings[0]
			uri = listOfStrings[1]
			firstLine++
		}
		if str == ""{
			fmt.Println(listOfStrings)
			response(conn,m,uri)
			break
		}
	}

}

func response(conn net.Conn,method string,uri string){

	fmt.Println(" I am in Response method : ",method, "URI : ",uri)

	switch method {
		case "GET":
			strToCmpare := "/"
			if strings.Compare(strToCmpare, uri) == 0 {
				body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
				<strong>INDEX</strong><br>
				<a href="/">index</a><br>
				<a href="/about">about</a><br>
				<a href="/contact">contact</a><br>
				<a href="/apply">apply</a><br>
				</body></html>`
				fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
				fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
				fmt.Fprint(conn, "Content-Type: text/html\r\n")
				fmt.Fprint(conn, "\r\n")
				fmt.Fprint(conn, body)

			}else{
				fmt.Fprint(conn, "HTTP/1.1 404 NOT_FOUND\r\n")
			}
		case "POST":
			strToCmpare := "/sothing/new"
			if strings.Compare(strToCmpare, uri) == 0 {
				body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
				<strong>APPLY PROCESS</strong><br>
				<a href="/">index</a><br>
				<a href="/about">about</a><br>
				<a href="/contact">contact</a><br>
				<a href="/apply">apply</a><br>
				</body></html>`

				fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
				fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
				fmt.Fprint(conn, "Content-Type: text/html\r\n")
				fmt.Fprint(conn, "\r\n")
				fmt.Fprint(conn, body)
			}else{
				fmt.Fprint(conn, "HTTP/1.1 400 BAD REQUEST\r\n")
			}

		case "PUT":
			strToCmpare := "/sothing/to/modify"
			if strings.Compare(strToCmpare, uri) == 0 {
				fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
				fmt.Fprint(conn, "Content-Type: text/html\r\n")
				fmt.Fprint(conn, "\r\n")
			}else{
				fmt.Fprint(conn, "HTTP/1.1 400 BAD REQUEST\r\n")
			}
		default:
			fmt.Fprintf(conn,"Not a valid case")
	}

	//wg.Done()

}
