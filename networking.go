package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

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
	go startAcceptingConnections(listen)
}

func startAcceptingConnections(listener net.Listener) {
	//should be in a loop because each batch (ie sendData) is a new Accept & Read
	for {
		conn, err := listener.Accept()
		CheckError(err)
		fmt.Println("listen.Accept")
		go handleIncomingConnection(conn)
	}
}

func handleIncomingConnection(conn net.Conn) {
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
