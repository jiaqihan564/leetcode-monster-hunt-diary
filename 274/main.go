package main

import (
	"fmt"
)

// func hIndex(citations []int) int {
// 	b := make([]int, 0)

// 	for k, v := range citations {
// 		if  {

// 		}
// 	}
// }

// hIndex 使用 O(n) 的计数桶算法计算 H 指数
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func hIndex(citations []int) int {
	n := len(citations)

	m := make([]int, n+1)
	c := 0

	for _, v := range citations {
		if v >= n {
			m[n]++
		} else {
			m[v]++
		}

	}

	for i := n; i >= 0; i-- {
		c += m[i]
		if c >= i {
			return i
		}
	}

	return c
}

func main() {
	tests := [][]int{
		{3, 0, 6, 1, 5},  // 答案: 3
		{1, 3, 1},        // 答案: 1
		{0, 0, 0},        // 答案: 0
		{10, 8, 5, 4, 3}, // 答案: 4
		{},               // 答案: 0
	}

	for _, t := range tests {
		fmt.Println(t, "=>", hIndex(t))
	}
}
