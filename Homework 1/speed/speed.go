package main

import (
	"fmt"
)

func main() {
	var speed float32

	fmt.Println("Speed converter")
	fmt.Println()

	fmt.Print("Enter speed (km/h): ")
	fmt.Scan(&speed)
	fmt.Println()

	fmt.Printf("Speed: %f m/s; %f km/min", speed/3.6, speed/60)
}
