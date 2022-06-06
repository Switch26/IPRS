package main

import (
	"fmt"
	"net/http"
)

// table
// API to extend table (os.Args)
// copy table to other node

var myTable map[string][]string // map (aka dictionary) of strings array (slice)
var skeleLinkGood = "https://skeles.s3.amazonaws.com/mp4/100000531923949209162209368389625844466268829635846690090339594494683882048735.mp4"
var skeleLinkBad = "https://skeles.s3.amazonaws.com/mp4/10000053192394920916220936838962584446626882963584669009033959449468388204873.mp4"

func main() {
	fmt.Println("main called")
	pin(skeleLinkGood)
}

func pin(link string) {
	//download link
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error: ", r)
		}
	}()
	downloadFile(link, "testination.mp4")

	//produce hash
	//add to myTable
	//delete downloaded file
}

func downloadFile(link, destFileName string) {
	// checking for errors in download
	resp, err := http.Get(link)
	if err != nil {
		panic(fmt.Sprintf("%v", err)) // no internet kinda deal
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("%v", resp)) // broken link
	}

	// by now we have a valid data
	fmt.Println(resp)

	//out, err := os.Create(destFileName)
	//if err != nil {
	//	panic(fmt.Sprintf("%v", err)) // can't create file
	//}
	//defer out.Close()
	//
	//n, err := io.Copy(out, resp.Body)
	//if err != nil {
	//	panic(fmt.Sprintf("%v", err)) // can't copy from buffer
	//}
	//fmt.Println("copied data: %v", n)
}
