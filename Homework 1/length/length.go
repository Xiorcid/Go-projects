package main

import (
	"fmt"
)

func main() {
	var a, b float32

	fmt.Println("Feets and inches converter")
	fmt.Println()

	fmt.Print("Enter feets: ")
	fmt.Scan(&a)
	fmt.Println()

	fmt.Print("Enter inches: ")
	fmt.Scan(&b)
	fmt.Println()

	fmt.Printf("Length: %fcm", a*30.48+b*2.54)
}
