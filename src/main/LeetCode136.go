package main

import "fmt"

func singleNumber(nums []int) int {
	ans := 0
	for _, tmp := range nums {
		ans = ans ^ tmp
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 2, 4}))
}
