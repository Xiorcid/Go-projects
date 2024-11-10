package main

import "fmt"

func main() {
	var nums [4]int16

	fmt.Print("Enter number line: ")
	fmt.Scan(&nums[0], &nums[1], &nums[2], &nums[3])

	for {
		numZero := nums[0]
		for i := 0; i <= 2; i++ {
			if nums[i] -= nums[i+1]; nums[i] < 0 {
				nums[i] = -nums[i]
			}
		}

		nums[3] -= numZero
		if nums[3] < 0 {
			nums[3] = -nums[3]
		}

		fmt.Println(nums[0], nums[1], nums[2], nums[3])

		if nums[0] == 0 && nums[1] == 0 && nums[2] == 0 && nums[3] == 0 {
			break
		}
	}
}
