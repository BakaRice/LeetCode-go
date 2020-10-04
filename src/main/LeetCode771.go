package main

/*
@Time : 2020/10/2 19:52
@Author : DELL ricemarch@foxmail.com
@tips: https://leetcode-cn.com/problems/jewels-and-stones/
*/

func numJewelsInStones(J string, S string) (sum int) {
	jewelsSet := map[byte]bool{}
	for i := range J {
		jewelsSet[J[i]] = true
	}
	for i := range S {
		if jewelsSet[S[i]] {
			sum++
		}
	}
	return
}
