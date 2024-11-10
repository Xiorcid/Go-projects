package main

import "fmt"

func main() {
	cells(10, 4)
	cells(5, 3)
	cells(3, 2)
}

func cells(n, m uint) {
	drawTop(m)
	var i uint
	for i = 1; i <= n; i++ {
		drawCell(m)
		if i+1 <= n {
			drawLine(m)
		}
	}
	drawBottom(m)
}

func drawTop(m uint) {
	fmt.Print("┌")
	var i uint
	for i = 1; i <= m; i++ {
		fmt.Print("─")
		if i+1 <= m {
			fmt.Print("┬")
		}
	}
	fmt.Println("┐")
}

func drawCell(m uint) {
	var i uint
	for i = 1; i <= m+1; i++ {
		fmt.Print("│ ")
	}
	fmt.Println()
}

func drawLine(m uint) {
	fmt.Print("├")
	var i uint
	for i = 1; i <= m; i++ {
		fmt.Print("─")
		if i+1 <= m {
			fmt.Print("┼")
		}
	}
	fmt.Println("┤")
}

func drawBottom(m uint) {
	fmt.Print("└")
	var i uint
	for i = 1; i <= m; i++ {
		fmt.Print("─")
		if i+1 <= m {
			fmt.Print("┴")
		}
	}
	fmt.Println("┘")
}
