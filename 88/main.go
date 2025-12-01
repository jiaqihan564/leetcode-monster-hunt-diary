package main

import "fmt"

// merge 函数将有序数组 nums2 合并到有序数组 nums1 中
// nums1 的长度为 m + n，其中前 m 个是有效有序元素，后 n 个是预留空间
// nums2 的长度为 n，且为有序
// 合并后结果要求依然有序，直接存放在 nums1 中（原地修改）
func merge(nums1 []int, m int, nums2 []int, n int) {
	// i 指向 nums1 有效元素的最后一个下标
	i := m - 1
	// j 指向 nums2 的最后一个下标
	j := n - 1
	// k 指向 nums1 的最后一个下标（整体长度为 m + n）
	k := m + n - 1

	// 只要 nums2 还有元素没有被合并完，就必须继续处理
	// 注意：循环条件是 j >= 0，而不是 i >= 0
	// 原因：nums1 前面的元素本身就已经在正确的位置上，没必要强行覆盖
	for j >= 0 {
		// 如果 i >= 0，说明 nums1 还有有效元素可以比较
		// 且当前 nums1[i] 比 nums2[j] 大，则把 nums1[i] 放到 nums1[k]
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			// 否则，把 nums2[j] 放到 nums1[k]
			// 包括两种情况：
			// 1. nums2[j] >= nums1[i]
			// 2. i < 0（nums1 有效区间已经全部用完）
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1) // 输出: [1 2 2 3 5 6]
}
