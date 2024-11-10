package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter A and B: ")
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	for i := a + 1; i < b; i++ {
		for j := 1; j <= i; j++ {
			if i == j*j {
				fmt.Println(i)
				break
			}
		}
	}
}
