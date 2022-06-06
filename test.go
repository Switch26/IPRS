package main

//The package name can be an arbitrary identifier, though if you want a package to serve as an entry point for an executable program, it needs to be named “main” and have a function main() with no arguments and no return type.

import "fmt"

//func main() {
//	f()
//	fmt.Println("Returned normally from f.")
//}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
