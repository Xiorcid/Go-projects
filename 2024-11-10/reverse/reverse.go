package main 

import "fmt"

func main(){
	a := []int{1, 2, 3, 4,5,6,7,8,9}
	Reverse(a)
	fmt.Print(a)
}

func Reverse(a []int){
	length := len(a)
	b:=make([]int, length)
	for i := range a{
		b[i] = a[length-i-1]
	}
	copy(a, b)
}