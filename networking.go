package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
)

func sendEcho(message string) {
	con, err := net.Dial("tcp", "djxmmx.net:17")
	CheckError(err)
	defer con.Close()

	_, err = con.Write([]byte(message))
	CheckError(err)

	//reply := make([]byte, 1024)
	//_, err = con.Read(reply)
	//CheckError(err)
	//fmt.Println(string(reply))
	//https://stackoverflow.com/questions/24339660/read-whole-data-with-golang-net-conn-read

	var myBuffer bytes.Buffer
	io.Copy(&myBuffer, con)
	fmt.Println(myBuffer.String())
}

func sendData(message, address string) {
	conn, err := net.Dial("tcp", address)
	CheckError(err)
	defer conn.Close()
	fmt.Fprintf(conn, message) // writing to connection
}

//every node should be constantly listening
func startListening(port string) {
	PORT := ":" + port
	listen, err := net.Listen("tcp", PORT) // opens port
	CheckError(err)
	//defer listen.Close()

	//for {
	//	conn, err := listen.Accept() // should it be in a loop?
	//	CheckError(err)
	//	//fmt.Println("listen.Accept")
	//	go handleIncomingConnection(conn)
	//}
	startAcceptingConnections(listen)
}

func startAcceptingConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept() // should it be in a loop?
		CheckError(err)
		fmt.Println("listen.Accept")
		go handleIncomingConnection(conn)
	}
}

func handleIncomingConnection(conn net.Conn) {
	//for {
	//	netData, err := bufio.NewReader(conn).ReadString('\n')
	//	fmt.Println("-> ", string(netData))
	//	CheckError(err)
	//}
	r := bufio.NewReader(conn)
	for {
		line, err := r.ReadBytes(byte('\n'))
		switch err {
		case nil:
			fmt.Println("->", string(line))
		case io.EOF:
		default:
			fmt.Println("ERROR: ", err)
		}

	}
}
