package main

import "fmt"

func main() {
	var a, b float32

	fmt.Println("Rectangle area calculator")
	fmt.Println()

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Printf("Area: %f; Perimeter: %f", a*b, 2*a+2*b)
}
