package main

//Problem: https://leetcode.com/problems/shuffle-the-array/submissions/
func shuffle(nums []int, n int) []int {

	var res []int
	for i, _ := range nums {
		if i+n >= len(nums) {
			break
		}
		x := nums[i]
		y := nums[i+n]
		res = append(res, x)
		res = append(res, y)
	}
	return res
}
