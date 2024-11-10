package main

import "fmt"


func main(){
	triangleLU(10)
	triangleLU(5)
	triangleLU(3)
}


func triangleLU(n uint){
	for ; n > 0; n--{
		var i uint
		for i = 0; i < n; i++{
			fmt.Print("X")
		}
		fmt.Println()
	}
}