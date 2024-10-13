package main

import "fmt"

func main() {
	var time, mins, hours, days int

	fmt.Print("Enter time (s): ")
	fmt.Scan(&time)

	days = time / 86400
	hours = (time - days*86400) / 3600
	mins = (time - days*86400 - hours*3600) / 60

	fmt.Printf("%d seconds = %d days, %d hours, %d minutes, %d seconds", time, days, hours, mins, time-days*86400-hours*3600-mins*60)
}
