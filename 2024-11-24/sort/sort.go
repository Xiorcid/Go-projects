package main

import(
	"fmt"
)


func main(){
	a := []int{5, 11, 7, 8, 5, 6, 1, 0, -1}
	b := CountSort(a)

	fmt.Printf("Array %v; Sorted %v", a, b)
}


func CountSort(a []int) []int{
	count, res := make([]int, len(a)), make([]int, len(a))

	for i, x := range a{
		for j := 0; j < len(a); j++{
			if a[j] < x || (a[j] == x && j<i){
				count[i]++
			}
		}
	}

	for i, x := range count{
		res[x] = a[i]
	}
	return res
}