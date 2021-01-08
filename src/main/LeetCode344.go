package main

/*
@Time : 2020/10/8 23:00
@Author : DELL ricemarch@foxmail.com
@tips:
*/

func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}
