package main

import (
	"fmt"
	"io"
	"os"
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

	fout, err := os.Create("starsv.res")
	if err != nil {
		return
	}
	defer fout.Close()

	ptr := 1 // Счетчик строки
	isMatched := false
	for ; ; ptr, isMatched = ptr+1, false {
		for _, x := range a { // Поставить *, если число меньше текущей строки, иначе пробел
			if x >= ptr {
				fmt.Fprint(fout, "*")
				isMatched = true
			} else {
				fmt.Fprint(fout, " ")
			}
		}
		fmt.Fprintln(fout)
		if !isMatched { // Если не выведена ни одна *, завершить
			break
		}
	}
}
