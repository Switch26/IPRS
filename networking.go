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

func startAcceptingConnecitons() {
	// this error checks for all panic error throws from methods below
	// this way I don't have to check error on every method
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("startAcceptingConnections() error: ", r)
		}
	}()

	myListener := openListener(1234)
	go runAcceptConnectionsLoop(myListener)

}

//every node should be constantly listening
func openListener(port int) net.Listener {
	PORT := ":" + string(port)
	listener, err := net.Listen("tcp", PORT) // opens port
	CheckError(err)
	return listener
	//defer listen.Close()
}

func runAcceptConnectionsLoop(listener net.Listener) {
	//should be in a loop because each batch (ie sendData) is a new Accept & Read
	for {
		conn, err := listener.Accept()
		CheckError(err)
		//fmt.Println("listen.Accept")
		handleIncomingConnection(conn)
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
