package main

import "fmt"

func main() {
	var x, pow float64
	fmt.Print("Enter X: ")
	fmt.Scan(&x)

	pow = x

	for i := 0; i < 99; i++ {
		pow = pow / (1 / x)
	}
	fmt.Printf("%f^100 = %f", x, pow)
}
