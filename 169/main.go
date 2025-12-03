package main

import "fmt"

func majorityElement(nums []int) int {
	var l int
	c := 0

	for _, v := range nums {
		if c == 0 {
			l = v
		}

		if l != v {
			c--
		} else {
			c++
		}

	}

	return l

}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(nums)) // 输出：2
}
