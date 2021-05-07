package main

import "fmt"

func main() {
	functionA(2, 3)    // 5
	functionB(2, 3)    // 6
	functionC(true)    // false
	functionD("$", 4)  // $ \n $ \n $ \n $ \n \n
	functionA(5, 6)    // 11
	functionB(5, 6)    // 30
	functionC(false)   // true
	functionD("ha", 3) // ha \n ha \n ha \n \n
}

func functionA(a int, b int) {
	fmt.Println(a + b)
}

func functionB(a int, b int) {
	fmt.Println(a * b)
}

func functionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Println(a)
	}
	fmt.Println()
}
