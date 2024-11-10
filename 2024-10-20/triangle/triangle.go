package main

import "fmt" 
import "math"

func main(){
	fmt.Println(triArea(3, 4, 5))
	fmt.Println(triArea(3, 1, 5))
	fmt.Println(triArea(-1, 4, 5))
	fmt.Println(triArea(30, 4, 5))
	fmt.Println(triArea(7, 8, 9))
	fmt.Println(triArea(-1, -4, -5))
}

func triArea(a, b, c float64) (area float64, correct bool){
	if a+b>c && a+c>b && b+c>a {
		p := (a+b+c)/2
		return math.Sqrt(p*(p-a)*(p-b)*(p-c)), true
	}
	return -1, false
}

