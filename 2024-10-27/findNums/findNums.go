package main

import "fmt"

func main(){
	getNums();
}


func isPrime(n int) bool{
	if(n%2 == 0 || n < 2){
		return false
	}
	i:=2
	for n%i != 0{i++}
	return n == i
}

func getNums(){
	n := 0
	for ;; {
		var i int
		for i = 1; i <= 50; i++{
			if isPrime(n+i){n += i; break}
		}
		if(i == 51){
			fmt.Printf("%d ... %d", n+1, n+50)
			break
		}
	}
}