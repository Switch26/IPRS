package main

import (
	"bytes"
	"fmt"
	cid "github.com/ipfs/go-cid"
	mc "github.com/multiformats/go-multicodec"
	mh "github.com/multiformats/go-multihash"
	"io"
	"net/http"
)

// table
// API to extend table (os.Args)
// copy table to other node
//cid "github.com/ipfs/go-cid"

var myTable map[string][]string // map (aka dictionary) of strings array (slice)
var skeleLinkGood = "https://skeles.s3.amazonaws.com/mp4/100000531923949209162209368389625844466268829635846690090339594494683882048735.mp4"
var skeleLinkBad = "https://skeles.s3.amazonaws.com/mp4/10000053192394920916220936838962584446626882963584669009033959449468388204873.mp4"

func main() {
	//fmt.Println("main called")
	//pin(skeleLinkGood)
	connectDatabase()
}

func pin(link string) {
	//download link
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error: ", r)
		}
	}()
	data := downloadFile(link, "testination.mp4")
	CIDstring := createCID(data)
	fmt.Println("Created CID: ", CIDstring)

	//add to myTable
	//delete downloaded file
}

func downloadFile(link, destFileName string) []byte {
	// checking for errors in download
	resp, err := http.Get(link)
	if err != nil {
		panic(fmt.Sprintf("%v", err)) // no internet kinda deal
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("%v", resp)) // broken link
	}
	//fmt.Println(resp)

	// by now we have valid data
	var myBuffer bytes.Buffer
	n, err := io.Copy(&myBuffer, resp.Body)
	if err != nil {
		panic(fmt.Sprintf("%v", err)) // can't copy from buffer
	}
	fmt.Println("copied data:", n)
	return myBuffer.Bytes()

	// save to a file
	//out, err := os.Create(destFileName)
	//if err != nil {
	//	panic(fmt.Sprintf("%v", err)) // can't create file
	//}
	//defer out.Close()
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
	if err != nil {
		println("error: %", err)
	}
	return c.String()
}
