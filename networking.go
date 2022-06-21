package main

import (
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

	resp, err := conn.Write([]byte(message))
	CheckError(err)
	fmt.Println("resp:", resp)
}
