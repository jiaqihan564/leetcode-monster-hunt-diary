package main

import "fmt"

func rotate(nums []int, k int) {
	m := len(nums)
	if m <= 1 {
		return
	}

	k = k % m
	if k == 0 {
		return
	}
	a(nums, 0, m-1)
	a(nums, 0, k-1)
	a(nums, k, m-1)

}

func a(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	// 构造一个示例数组
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println("旋转前：", nums)
	rotate(nums, k)
	fmt.Println("旋转后：", nums)
}
