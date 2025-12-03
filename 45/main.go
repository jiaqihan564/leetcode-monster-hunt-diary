package main

import "fmt"

func jump(nums []int) int {

	if len(nums) < 2 {
		return 0
	}

	n := 0
	m := 0

	o := 0

	for k, v := range nums {
		if k+v > m {
			m = k + v
		}

		if n == k {
			o++
			n = m
			if n >= len(nums)-1 {
				return o
			}
		}
	}

	return o
}

// jump 是力扣第 45 题的最优解：时间复杂度 O(n)，空间复杂度 O(1)
// 思路：单次遍历 + 贪心，按“层级 BFS”的方式维护当前步数能覆盖的区间
func jump1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		// 只有一个元素（已经在终点），不需要跳
		return 0
	}

	// steps      ：当前已经使用的跳跃次数
	// currentEnd ：当前这一步（当前层）所能覆盖的最远下标
	// furthest   ：在扫描当前层的过程中，下一步（下一层）能到达的最远下标
	steps := 0
	currentEnd := 0
	furthest := 0

	// 只需要遍历到 n-2：
	// 当我们走到最后一个下标 n-1 时，其实已经不需要再跳了
	for i := 0; i < n-1; i++ {
		// 更新从当前层任一点 i 起跳时，下一步能到达的最远位置
		if i+nums[i] > furthest {
			furthest = i + nums[i]
		}

		// 当 i 走到了当前层的最右边界 currentEnd：
		// 说明当前层的所有位置都看完了，必须要进行一次“跳跃”进入下一层
		if i == currentEnd {
			steps++               // 消耗一次跳跃
			currentEnd = furthest // 更新下一层的边界

			// 小优化：如果当前层已经可以覆盖到或超过终点，则可以提前结束
			if currentEnd >= n-1 {
				break
			}
		}
	}

	return steps
}

func main() {
	// 简单测试用例
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 输出 2，路径：0 -> 1 -> 4
	fmt.Println(jump([]int{2, 3, 0, 1, 4})) // 输出 2
	fmt.Println(jump([]int{1, 2, 1, 1, 1})) // 输出 3
}
