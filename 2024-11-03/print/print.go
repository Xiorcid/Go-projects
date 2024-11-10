package main

import "fmt"

type myArr [15]int

func main(){
	a := myArr{2, 3, -3, 7, -6, 0, 2, 5, -4, -9, 1, -5, 8, 5, 6}
	Print1(a)
	Print2(a)
}

func Print1(a myArr){
	for _, x := range a{
		if(x>0){
			fmt.Println(x)
		}else{
			fmt.Print("  ")
			fmt.Println(x)
		}
	}

	fmt.Println()
}

func Print2(a myArr){
	var pos, neg myArr
	var posptr, negptr int
	for _, x := range a{
		if(x>0){
			pos[posptr] = x
			posptr++
		}else{
			neg[negptr] = x
			negptr++
		}
	}

	for i := 0; ; i++{
		if posptr>i {
			fmt.Printf("%d ", pos[i]) 
		}else{
			fmt.Print("  ")
		}
		if negptr>i{fmt.Printf("%d ", neg[i])}
		fmt.Println()
		if posptr < i && negptr < i {break}
	}
}