package main

import "fmt"

func main() {
	for i := 1000; i <= 10000; i++ {
		if i%7 == 0 && (i%10+(i%100-i%10)/10+(i%1000-i%100)/100+(i%10000-i%1000)/1000)%5 == 0 {
			fmt.Println(i)
		}
	}
}
