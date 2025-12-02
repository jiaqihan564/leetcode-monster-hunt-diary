package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums) - 1

	for i := 0; i <= n; {
		if nums[i] == i {
			nums[i] = nums[n]
			n--
		} else {
			i++
		}
	}
	return n
}

// removeElement 移除切片 nums 中所有等于 val 的元素
// 返回值为移除后数组的“新长度”
// 要求：在原切片上原地修改，且额外空间为 O(1)
// 本实现采用“头尾双指针交换”策略，在 val 较少时能减少写操作次数。
func removeElement1(nums []int, val int) int {
	// n 表示当前“有效区间”的长度，初始为整个切片长度
	n := len(nums)

	// i 从前往后扫描，只要 i < n 就说明还在有效区间里
	for i := 0; i < n; {
		if nums[i] == val {
			// 如果当前元素等于要删除的值 val：
			// 用有效区间最后一个元素覆盖当前元素
			// 相当于：把要删除的元素“丢到数组末尾”，并且缩短有效区间
			nums[i] = nums[n-1]
			n-- // 有效区间长度减 1

			// 注意：这里不能 i++ !!!
			// 因为 nums[i] 被替换成了从尾部拿来的新元素，
			// 这个新元素还没有检查过是否等于 val，
			// 所以下一轮循环还需要继续检查 nums[i]。
		} else {
			// 如果当前元素不等于 val，则它是“保留元素”，i 往后移动即可
			i++
		}
	}

	// 此时，nums[0:n] 中存放的就是删除 val 后剩下的所有元素（顺序不保证）
	// nums[n:] 部分的内容无所谓，题目不会查看。
	return n
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElement(nums, val)
	fmt.Println("new length:", k, "nums:", nums[:k]) // 输出: new length: 2 nums: [2 2]

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	k2 := removeElement1(nums2, val2)
	fmt.Println("new length:", k2, "nums2:", nums2[:k2]) // 可能是 [0 1 4 0 3] 等
}
