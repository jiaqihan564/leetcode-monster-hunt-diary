package main

func merge1(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if nums1[j] > nums2[i] {
			nums1[k] = nums2[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

}
