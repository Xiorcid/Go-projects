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

	fout, err := os.Create("stars.res")
	if err != nil{
		return
	}
	defer fout.Close()
	
	var c int
	for {
		if _, err = fmt.Fscanf(fin, "%d\n", &c); err == io.EOF {
			break
		}
		for i := 0; i < c; i++{
			fmt.Fprint(fout, "*")
		}
		fmt.Fprintln(fout)
	}
}
