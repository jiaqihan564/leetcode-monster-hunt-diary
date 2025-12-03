package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	l := 2

	for k := 2; k < n; k++ {
		if nums[k] != nums[l-2] {
			nums[l] = nums[k]
			l++
		}
	}

	return l
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	fmt.Println("new length:", k)
	fmt.Println("new array:", nums[:k])
}
