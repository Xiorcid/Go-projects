package main

import "fmt"

func main() {
	fmt.Println("Determinant 3*3 calc")
	fmt.Println()

	var matrix [3][3]int

	for i := 0; i < 3; i++ {
		fmt.Printf("Enter %d row: ", i+1)
		fmt.Scan(&matrix[i][0], &matrix[i][1], &matrix[i][2])
	}
	fmt.Println()

	det := (matrix[0][0] * matrix[1][1] * matrix[2][2]) + (matrix[1][0] * matrix[2][1] * matrix[0][2]) + (matrix[0][1] * matrix[1][2] * matrix[2][0])
	det -= (matrix[0][2] * matrix[1][1] * matrix[2][0]) + (matrix[1][2] * matrix[2][1] * matrix[0][0]) + (matrix[0][1] * matrix[1][0] * matrix[2][2])
	// 00 01 02
	// 10 11 12
	// 20 21 22

	fmt.Printf("Determinant: %d", det)

}
