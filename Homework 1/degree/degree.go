package main

import (
	"fmt"
	"math"
)

func main() {
	var x int32

	fmt.Println("Degree calculator")
	fmt.Println()

	fmt.Print("Enter x: ")
	fmt.Scan(&x)
	fmt.Println()

	for i := 2; i < 6; i++ {
		fmt.Printf("%d^%d: %d; ", x, i, int((math.Pow(float64(x), float64(i)))))
	}

}
