package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileinfo, err := os.Stat("my.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fileinfo.Size())
	}
}