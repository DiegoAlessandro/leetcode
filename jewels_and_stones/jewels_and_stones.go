package main

// https://leetcode.com/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {

	cnt := 0
	jewelsType := jewels
	for _, t := range jewelsType {
		for _, stone := range stones {
			if t == stone {
				cnt = cnt + 1
			}
		}
	}
	return cnt
}
