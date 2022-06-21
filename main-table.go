package main

import (
	"bufio"
	"bytes"
	"fmt"
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

// API to extend table (os.Args)
// copy table to other node

var skeleLinkGood = "https://ipfs.io/ipfs/QmV6aq5mm82YtoDGinPwDVrNWAyNegaJJYtW7kj1NdZpdB"

func main() {
	println("main called")
	//newCID, err := pin(skeleLinkGood)
	//println("result: ", newCID, err)
	//sendEcho("yo!")

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}

func pin(link string) (string, error) {
	// just handling panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error: ", r)
		}
	}()

	data := downloadFile(link, "destination.mp4")
	CIDstring := createCID(data)
	fmt.Println("Created CID: ", CIDstring)
	_, err := insertRow(CIDstring, link)
	if err == nil {
		return CIDstring, nil
	} else {
		return "", err
	}
}

func downloadFile(link, destFileName string) []byte {
	// checking for errors in download
	resp, err := http.Get(link)
	CheckError(err)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("%v", resp)) // broken link
	}
	//fmt.Println(resp)

	// by now we have valid data
	var myBuffer bytes.Buffer
	n, err := io.Copy(&myBuffer, resp.Body)
	CheckError(err)
	fmt.Println("copied data:", n)
	return myBuffer.Bytes()
}

func createCID(data []byte) string {
	//println(uint64(mc.Raw))
	pref := cid.Prefix{
		Version:  1,
		Codec:    uint64(mc.DagPb),
		MhType:   mh.SHA2_256,
		MhLength: -1, // default length
	}
	c, err := pref.Sum([]byte(data))
	CheckError(err)
	return c.String()
}

func CheckError(err error) {
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

// save to a file
//out, err := os.Create(destFileName)
//if err != nil {
//	panic(fmt.Sprintf("%v", err)) // can't create file
//}
//defer out.Close()
