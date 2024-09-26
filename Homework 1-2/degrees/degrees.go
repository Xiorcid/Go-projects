package main

import (
	"fmt"
)

func main() {
	var degrees float64

	fmt.Println("Degrees converter")
	fmt.Println()

	fmt.Print("Enter value (C°): ")
	fmt.Scan(&degrees)
	fmt.Println()

	fmt.Printf("R=%f° F=%f° K=%f°", degrees*0.8, degrees*1.8+32, degrees+273.15)
}
