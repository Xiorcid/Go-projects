package main

import "fmt"

func main(){
	fmt.Println(msd(25417792)) // 2
	fmt.Println(msd(4567)) // 2
	fmt.Println(msd(0)) // 0
	fmt.Println(msd(10)) // 1
}

func msd(n uint) uint {
	var x uint
	for i := 1; pow10(i-1) <= n; i++{
		if pow10(i-1) == 0{continue}
		if n / pow10(i) == 0{	
			x = n / pow10(i-1)
			break
		}
	}
	return x
}

func pow10(x int) uint{
	var i int
	var n uint
	n = 1
	for i = 0; i < x; i++{
		n *= 10
	}
	return n
}