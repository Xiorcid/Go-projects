package main

import "fmt"

func main() {
	var height, width float64

	fmt.Print("Enter width: ")
	fmt.Scan(&width)
	fmt.Println()

	fmt.Print("Enter height: ")
	fmt.Scan(&height)
	fmt.Println()

	fmt.Println("Area: ", width*height, "; perimeter: ", 2*width+2*height)
}
