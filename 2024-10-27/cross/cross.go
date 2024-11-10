package main

import "fmt"


func main(){
	cross(9)
	cross(4)
	cross(3)
}


func cross(n uint) bool{
	if (n%2!=0 && n > 0){
		var i uint
		var flag bool
		for i = 1; ;{
			drawSpace(i)
			fmt.Print("X")
			drawSpace(n-2*i+1)
			if((n-2*i+1) == 0){
				fmt.Println(); 
				i--; 
				flag = true; 
				continue
			}
			fmt.Print("X")
			fmt.Println()
			if flag{
				i--
			}else{
				i++
			}
			if(i == 0){return true}
		}
	}
	return false
}


func drawSpace(n uint){
	if(n > 1){
		n--
		for ; n > 0; n--{
			fmt.Print(" ")
		}
	}
}