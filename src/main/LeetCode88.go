package main

import "fmt"

/*
@Time : 2020/8/12 15:27
@Author : DELL ricemarch@foxmail.com
@tips: https://leetcode-cn.com/problems/merge-sorted-array/
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	max := m + n
	left, right := m-1, n-1
	i := 1
	for left >= 0 && right >= 0 {
		if nums1[left] < nums2[right] {
			nums1[max-i] = nums2[right]
			i++
			right--
		} else {
			nums1[max-i] = nums1[left]
			i++
			left--
		}
	}
	if left == -1 {
		for i := right; i >= 0; i-- {
			nums1[i] = nums2[i]
		}
	}
}

func main() {
	var nums1 = []int{0}
	merge(nums1, 0, []int{1}, 1)
	fmt.Println(nums1)
}
