package main

import "fmt"

func main(){
	a := []int{1, 0, 3, 0, 5, 6, 0, 8, 9, 0, 10, 11}
	fmt.Println(Purge0(a))
}

func Purge0(a []int) []int{
	var startptr, endptr int
	b := make([] int, 0)
	for i, x := range a{
		if(endptr > startptr){startptr = i}
		if(x == 0 || i == len(a)-1){
			endptr = i;
			if(i == len(a)-1){endptr+=1}
			b = append(b, a[startptr:endptr]...)
			fmt.Println(a[startptr:endptr], startptr, endptr)
		}
	}
	return b 
}