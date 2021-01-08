package main

import "sort"

/*
@Time : 2020/10/16 9:13
@Author : DELL ricemarch@foxmail.com
@tips:
*/
func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, v := range A {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}
