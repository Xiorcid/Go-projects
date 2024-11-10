package main

import "fmt"

type myArr [15]int

func main(){
	fmt.Println(fillA())
	fmt.Println(fillB())
	fmt.Println(fillC())
	fmt.Println(fillD())
	fmt.Println(fillE())
}

func fillA() myArr{
	var arr myArr
	for i := 0; i < len(arr); i++{
		if i%2 == 1{
			arr[i] = 1
		}
	}
	return arr
}

func fillB() myArr{
	var arr myArr
	cnt := 1
	for i := 0; i < len(arr); i++{
		if i%2 == 0{
			arr[i] = cnt
			cnt++
		}
	}
	return arr
}

func fillC() myArr{
	var arr myArr
	cnt, ptr := 1, 1
	for i := 0; i < len(arr); i++{
		arr[i] = cnt
		ptr--
		if ptr == 0 {
			cnt++
			ptr = cnt
		}
	}
	return arr
}

func fillD() myArr{
	var arr myArr
	ctr := 1
	for i := 0; i < len(arr); i+=ctr{
		arr[i] = 1
		ctr++
	}
	return arr
}

func fillE() myArr{
	var arr myArr
	ctr := 2
	arr[0] = 1
	for i := 1; i < len(arr); i++{
		arr[i] = arr[i-1]+ctr 
		ctr+=2
	}
	return arr
}