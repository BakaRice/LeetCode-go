package main

/*
@Time : 2020/8/9 17:31
@Author : DELL ricemarch@foxmail.com
@tips: https://leetcode-cn.com/problems/search-a-2d-matrix-ii/comments/
*/

func searchMatrix(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0 {
		return false
	}
	shorterDim := Min(len(matrix), len(matrix[0]))
	for i := 0; i < shorterDim; i++ {
		verticalFound := binarySearch(matrix, target, i, true)
		horizontalFound := binarySearch(matrix, target, i, false)
		if verticalFound || horizontalFound {
			return true
		}
	}
	return false
}

func binarySearch(matrix [][]int, target int, start int, vertical bool) bool {
	lo := start
	var hi int
	if vertical {
		hi = len(matrix[0]) - 1
	} else {
		hi = len(matrix) - 1
	}
	for {
		if hi < lo {
			break
		}
		mid := (lo + hi) / 2
		if vertical {
			if matrix[start][mid] < target {
				lo = mid + 1
			} else if matrix[start][mid] > target {
				hi = mid - 1
			} else {
				return true
			}
		} else {
			if matrix[mid][start] < target {
				lo = mid + 1
			} else if matrix[mid][start] > target {
				hi = mid - 1
			} else {
				return true
			}
		}
	}
	return false
}
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
