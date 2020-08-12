package main

import "fmt"

/*
@Time : 2020/8/12 20:32
@Author : DELL ricemarch@foxmail.com
@tips:
*/

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	maxlen := 0
	for _, v := range wordDict {
		set[v] = true
		if len(v) > maxlen {
			maxlen = len(v)
		}
	}
	dp := make([]int, len(s))
	j := 0
	for i := j; i <= len(s); i++ {
		subStr := s[j:i]
		if set[subStr] {
			dp[i-1] = 1
			if dp[len(s)-1] == 1 {
				return true
			}
		}
		if i-j == maxlen {
			for p := j; p < i; p++ {
				if dp[p] == 1 {
					j = p + 1
					break
				}
			}
		}
	}
	return false
}

//抄 官方题解
//dp[i]=dp[j] && check(s[j..i−1])
//其中 check(s[j..i−1])
//表示子串 s[j..i-1]s[j..i−1] 是否出现在字典中。
func wordBreak_2(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	var bool = wordBreak_2("leetcode", []string{"leet", "code"})
	var bool2 = wordBreak_2("aaaaaaa", []string{"aaaa", "aaa"})
	var bool3 = wordBreak_2("cars", []string{"car", "ca", "rs"})
	var bool4 = wordBreak_2("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	var bool5 = wordBreak_2("a", []string{"b"})
	var bool6 = wordBreak_2("bb", []string{"a", "b", "bbb", "bbbb"})

	fmt.Println(bool)
	fmt.Println(bool2)
	fmt.Println(bool3)
	fmt.Println(bool4)
	fmt.Println(bool5)
	fmt.Println(bool6)
}
