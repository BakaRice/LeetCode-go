package main

/*
@Time : 2020/10/11 13:02
@Author : DELL ricemarch@foxmail.com
@tips:
*/

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	if max > target {
		return false
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	dp[0][nums[0]] = true
	for i := 1; i < n; i++ {
		v := nums[i]
		for j := 1; j <= target; j++ {
			if j >= v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n-1][target]
}
