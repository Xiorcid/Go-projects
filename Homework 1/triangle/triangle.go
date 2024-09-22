package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p float64

	fmt.Println("Triangle area calculator")
	fmt.Println()

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Print("Enter c: ")
	fmt.Scan(&c)
	fmt.Println()

	p = (a + b + c) / 2

	fmt.Printf("Area: %f", math.Sqrt(p*(p-a)*(p-b)*(p-c)))
}
