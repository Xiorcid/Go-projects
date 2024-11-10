package main

import (
	"fmt"
	"math"
)

func main() {
	var speed float64
	fmt.Print("Enter speed (km/h): ")
	fmt.Scan(&speed)
	fmt.Println()

	fmt.Print("Speed: ", math.Round(speed/3.6), " m/s, ", math.Round(speed/60), " km/m")
}
