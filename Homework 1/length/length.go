package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b string

	fmt.Println("Feets and inches converter")
	fmt.Println()

	fmt.Print("Enter length (3ft 5in): ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println()

	ft, err := strconv.Atoi(strings.Split(a, "ft")[0])
	in, err := strconv.Atoi(strings.Split(b, "in")[0])
	if err == nil {
		fmt.Printf("Length: %fcm", float32(ft)*30.48+float32(in)*2.54)
	}
}
