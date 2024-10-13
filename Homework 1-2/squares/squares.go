package main

import "fmt"

func main() {
	var x1, x2 int

	fmt.Println("Get square equation")
	fmt.Println()

	fmt.Print("Enter 1st square: ")
	fmt.Scan(&x1)
	fmt.Println()

	fmt.Print("Enter 2st square: ")
	fmt.Scan(&x2)
	fmt.Println()

	fmt.Printf("Square equation: X^2+%dX+%d=0", ((x1 + x2) * -1), x1*x2)
}
