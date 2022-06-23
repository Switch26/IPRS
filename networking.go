package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
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

	resp, err := conn.Write([]byte(message))
	CheckError(err)
	fmt.Println("resp:", resp)
}

//every node should be constantly listening
func startListening(port string) {
	listen, err := net.Listen("tcp", port) // opens port
	CheckError(err)
	defer listen.Close()

	for {
		conn, err := listen.Accept() // should it be in a loop?
		CheckError(err)
		go handleIncomingConnection(conn)
	}
}

func handleIncomingConnection(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(3 * time.Minute))
	defer conn.Close()

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		CheckError(err)
		fmt.Println("-> ", string(netData))
	}
}
