package main

import "fmt"

func main() {
	var amount, total float64
	amount = getPaintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f litters needed.\n", amount)
	total += amount
	amount = getPaintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f litters needed.\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f litters\n", total)

}

func getPaintNeeded(a float64, b float64) float64 {
	return a * b / 10.0
}
