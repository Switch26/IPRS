package main

//import (
//	"bufio"
//	"fmt"
//	"net"
//	"os"
//	"strings"
//	"time"
//)
//
//func main() {
//	arguments := os.Args
//	if len(arguments) == 1 {
//		fmt.Println("Please provide port number")
//		return
//	}
//
//	PORT := ":" + arguments[1]
//	listen, err := net.Listen("tcp", PORT)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer listen.Close()
//
//	conn, err := listen.Accept()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for {
//		netData, err := bufio.NewReader(conn).ReadString('\n')
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		if strings.TrimSpace(string(netData)) == "STOP" {
//			fmt.Println("Exiting TCP server")
//			return
//		}
//
//		fmt.Print("-> ", string(netData))
//		t := time.Now()
//		myTime := t.Format(time.RFC3339) + "\n"
//		conn.Write([]byte(myTime))
//	}
//}
