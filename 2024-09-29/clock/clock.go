package main

import "fmt"

func main() {
	var hours, mins float32

	fmt.Print("Enter hours: ")
	fmt.Scan(&hours)
	fmt.Print("Enter minutes: ")
	fmt.Scan(&mins)

	fmt.Printf("Hours hand angle: %f, minutes hand angle: %f", 30*hours, 60*mins)
}
