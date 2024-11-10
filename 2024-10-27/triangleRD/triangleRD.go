package main

import "fmt"


func main(){
	triangleRD(10)
	triangleRD(5)
	triangleRD(3)
}


func triangleRD(n uint){
	w := n
	for ; n > 0; n--{
		drawSpace(n)
		var i uint
		for i = 0; i+n <= w; i++{
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