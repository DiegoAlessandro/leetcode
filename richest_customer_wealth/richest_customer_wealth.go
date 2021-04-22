package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {

	var sumAmount []int
	for _, customer := range accounts {
		var sum = 0
		for _, amount := range customer {
			sum = amount + sum
		}
		sumAmount = append(sumAmount, sum)
	}

	max := -999999
	for _, amount := range sumAmount {
		if amount > max {
			max = amount
		}
	}

	return max
}

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	res := maximumWealth(accounts)
	fmt.Println(res)
}
