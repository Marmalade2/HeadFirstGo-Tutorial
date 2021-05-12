package main

import "fmt"

func main() {
	myInt, myBool, myString := many_return()
	fmt.Println(myInt, myBool, myString)
}

func many_return() (int, bool, string) {
	return 1, false, "Hello"
}
