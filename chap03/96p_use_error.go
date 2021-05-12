package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// error value => That is a value what has Error method which return string
	err := errors.New("This is new error.") // create new error
	fmt.Println(err.Error())                // Error return string
	fmt.Println(err)                        // fmt package print error string automatically if parameter has Error() method.
	log.Println(err)                        // log also print error string in same condition.

	//if you want to set formatting error message, use Errorf in fmt package.
	err2 := fmt.Errorf("%.2f is negative value.", -2.3894)
	fmt.Println(err2)
	log.Fatal(err2)
}
