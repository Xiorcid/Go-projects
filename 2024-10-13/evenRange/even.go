package main

import "fmt"

func main() {
	var a, b int16

	fmt.Print("Enter A and B: ")
	fmt.Scan(&a, &b)

	if a > b {
		a, b = b, a
	}

	for i := a + 1; i < b; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
