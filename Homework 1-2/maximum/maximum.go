package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Find maximal number")
	fmt.Println()

	fmt.Print("Enter 1st number: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Enter 2st number: ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Printf("Maximum: %f", math.Max(a, b))
}
