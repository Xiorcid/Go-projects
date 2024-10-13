package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Enter 1st number: ")
	fmt.Scan(&a)
	fmt.Print("Enter 2st number: ")
	fmt.Scan(&b)
	fmt.Print("Enter 3st number: ")
	fmt.Scan(&c)

	var out_num [3]int
	if a > b && a > c {
		out_num[2] = a
	} else if b > a && b > c {
		out_num[2] = b
	} else if c > a && c > b {
		out_num[2] = c
	}
	if a > b && a > c {
		out_num[1] = a
	} else if b > a && b > c {
		out_num[2] = b
	} else if c > a && c > b {
		out_num[2] = c
	}
	if a > b && a > c {
		out_num[1] = a
	} else if b < a && b < c {
		out_num[1] = b
	} else if c < a && c < b {
		out_num[1] = c
	}

	fmt.Print(out_num[0], ", ", out_num[1], ", ", out_num[2])
}
