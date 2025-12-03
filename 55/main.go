package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)

	if n < 2 {
		return true
	}

	y := 0

	for k, v := range nums {
		if k > y {
			return false
		}

		if v+k > y {
			y = v + k
		}

		if y >= n-1 {
			return true
		}
	}

	return false
}

func main() {
	tests := [][]int{
		{2, 3, 1, 1, 4}, // true
		{3, 2, 1, 0, 4}, // false
		{0},             // true
		{1, 0, 1, 0},    // false
		{2, 0},          // true
	}

	for _, nums := range tests {
		fmt.Println(nums, "=>", canJump(nums))
	}
}
