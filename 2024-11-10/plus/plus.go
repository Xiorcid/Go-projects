package main

import "fmt"

func main() {
	a := []int{2, 3, 0, 1, 5, 8}
	Plus(a) // 2, 5, 3, 1, 6, 13
	fmt.Println(a)
}

func Plus(a []int) {
	for i := range a {
		if i == 0 {
			continue
		}
		a[i] += a[i-1]
	}
}
