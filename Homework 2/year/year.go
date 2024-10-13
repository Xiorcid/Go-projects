package main

import "fmt"

func main() {
	var year uint16
	isLeap := false
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if year%4 == 0 {
		isLeap = true
		if year%100 == 0 {
			if year%400 != 0 {
				isLeap = false
			}
		}
	}

	if isLeap {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d isn't a leap year", year)
	}
}
