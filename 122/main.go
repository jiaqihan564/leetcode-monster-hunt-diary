package main

import "fmt"

func maxProfit(prices []int) int {
	m := len(prices)
	if m < 2 {
		return 0
	}

	max := 0

	for i := 1; i < m; i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}

	return max
}

func main() {
	// 一些测试用例，方便本地快速验证逻辑
	tests := [][]int{
		{7, 1, 5, 3, 6, 4}, // 1->5 赚4，3->6 赚3，总共 7
		{1, 2, 3, 4, 5},    // 每天都涨，(1->2)+(2->3)+(3->4)+(4->5)=4
		{7, 6, 4, 3, 1},    // 一路下跌，根本不交易，利润 0
		{2, 1, 2, 0, 1},    // (1->2)=1, (0->1)=1，总共 2
	}

	for _, prices := range tests {
		fmt.Println(prices, "=>", maxProfit(prices))
	}
}
