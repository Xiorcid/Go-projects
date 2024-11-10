package main

import "fmt"

func main(){
	fmt.Println(Power2(65536))
}

func Power2(max int64) []int64{
	var st int64
	a := make([]int64, 0)
	for i := 1; ; i++{
		st = pow(2, i)
		if(st>max){return a}
		a = append(a, st)
	}
}

func pow(num int64, st int) int64{
	x:=num
	for i := 1; i < st; i++{
		num *= x
	}
	return num
}