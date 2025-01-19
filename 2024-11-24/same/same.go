package main 

import(
	"bufio"
	"os"
	"fmt"
	"io"
)

func main(){
	fin_1, err := os.Open("n1.in")
	if err != nil{
		return
	}
	defer fin_1.Close()
	fin_1_buf := bufio.NewReader(fin_1)

	fin_2, err := os.Open("n2.in")
	if err != nil{
		return
	}
	defer fin_2.Close()
	fin_2_buf := bufio.NewReader(fin_2)


}