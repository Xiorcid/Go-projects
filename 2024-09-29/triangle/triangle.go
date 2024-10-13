package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter 1st : ")
	fmt.Scan(&a)
	fmt.Print("Enter 2st : ")
	fmt.Scan(&b)
	fmt.Print("Enter 3st : ")
	fmt.Scan(&c)

	if a+b > c && a+c > b && c+b > a {
		fmt.Printf("This triangle exists")
	} else {
		fmt.Printf("Incorrect triangle")
	}
}
