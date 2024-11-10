package main

import "fmt"

func main(){
	var n, i, oldD uint
	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	for i = 0; i <= n; i++{
		if d := nextDivision(n, i); d != 0 && oldD != d{
			fmt.Println(d)
			oldD = d
		}
	}	
}

func nextDivision(n uint, d uint) uint{
	for d++; d <= n; d++{
		if n % d == 0{
			return d
		}
	}
	return 0
}