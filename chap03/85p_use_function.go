package main

import "fmt"

func main() {
	printPaintLitters(3.2, 4.8)
	printPaintLitters(5.7, 9.5)
	fmt.Println(getPaintLitters(10.5, 9.7))
	fmt.Printf("%.2f\n", getPaintLitters(4.13, 9.59))
}

func printPaintLitters(width float64, height float64) {
	fmt.Printf("%.2f liters needed.\n", width*height/10.0)
}
func getPaintLitters(width float64, height float64) float64 {
	return width * height / 10.0
}
