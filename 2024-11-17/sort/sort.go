package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	fin, err := os.Open("numbers.in")
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	defer fin.Close()

	var a []int
	for {
		var c int
		if _, err = fmt.Fscanf(fin, "%d\n", &c); err == io.EOF {
			break
		}
		a = append(a, c)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	fmt.Print(a[0:5])
}
