package main

import "fmt"

func removeDuplicates(nums []int) int {
	a := len(nums)
	if a <= 1 {
		return a
	}

	m := 0

	for i := 1; i < a; i++ {
		if nums[i] != nums[m] {
			m++
			nums[m] = nums[i]
		}
	}

	return m + 1

}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)

	fmt.Println("去重后长度:", k)
	fmt.Println("去重后的前 k 个元素:", nums[:k])
}
