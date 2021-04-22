package main

// Problem: https://leetcode.com/problems/number-of-good-pairs/
func numIdenticalPairs(nums []int) int {

	var cnt = 0
	for i, _ := range nums {

		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				cnt += 1
			}
		}
	}
	return cnt
}
