package main

import "math"

/*
@Time : 2020/10/14 10:49
@Author : DELL ricemarch@foxmail.com
@tips:
*/

func commonChars(A []string) []string {
	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range A {
		freq := [26]int{}
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq[:] {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans
}
