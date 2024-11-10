package main

import "fmt"

func main(){
	fmt.Println(reverse(1234))
	fmt.Println(reverse(8956))
	fmt.Println(reverse(1000))
}

func reverse(n uint) uint{
	var res uint

	return res
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

// if uint(math.Pow10(i-1)) == 0{continue}
// fmt.Println( n / uint(math.Pow10(i-1)))
// res += n / uint(math.Pow10(i-1))