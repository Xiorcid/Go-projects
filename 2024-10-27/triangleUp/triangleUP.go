package main

import "fmt"


func main(){
	triangleUP(10)
	triangleUP(5)
	triangleUP(3)
}


func triangleUP(n uint){
	last := 2*n-1
	var i uint
	for i = 1; i <= last; i+=2{
		drawSpace(uint((last-1)/2)-(i/2-1))
		var j uint
		for j = 0; j < i; j++{
			fmt.Print("X")
		}
		fmt.Println()
	}
}

func drawSpace(n uint){
	n--
	for ; n > 0; n--{
		fmt.Print(" ")
	}
}