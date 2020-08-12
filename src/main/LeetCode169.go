package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	hp := map[int]int{}
	n2 := len(nums) / 2
	for _, v := range nums {
		hp[v]++
		if hp[v] > n2 {
			return v
		}
	}
	return 0
}

func majorityElement_2(nums []int) int {
	cand_num := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if cand_num == nums[i] {
			count = count + 1
		} else {
			count = count - 1
			if count == 0 {
				cand_num = nums[i]
				count = 1
			}
		}
	}
	return cand_num
}

func main() {
	fmt.Println(majorityElement_2([]int{8, 8, 7, 7, 7}))
}
