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

	fout0, err := os.Create("numbers0.res")
	if err != nil {
		return
	}
	fout1, err := os.Create("numbers1.res")
	if err != nil {
		return
	}
	defer fout0.Close()
	defer fout1.Close()

	var c int
	for {
		if _, err = fmt.Fscanf(fin, "%d\n", &c); err == io.EOF {
			break
		}
		if c%2 == 0 {
			fmt.Fprintln(fout0, c)
		} else {
			fmt.Fprintln(fout1, c)
		}
	}
}
