package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	m := len(prices)
	if m < 2 {
		return 0
	}

	min := math.MaxInt

	max := 0

	for _, v := range prices {
		if v < min {
			min = v
		}

		l := v - min

		if l > max {
			max = l
		}

	}

	return max
}

func main() {
	// 一些简单用例验证
	tests := [][]int{
		{7, 1, 5, 3, 6, 4}, // 预期：5（1 买入，6 卖出）
		{7, 6, 4, 3, 1},    // 预期：0（一直在跌，不买）
		{1, 2, 3, 4, 5},    // 预期：4（1 买入，5 卖出）
		{2, 4, 1},          // 预期：2（2 买入，4 卖出）
	}

	for _, prices := range tests {
		fmt.Println(prices, "=>", maxProfit(prices))
	}
}
