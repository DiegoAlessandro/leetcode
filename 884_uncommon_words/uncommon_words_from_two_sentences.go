package main

import (
	"strings"
)

// https://leetcode.com/problems/uncommon-words-from-two-sentences/
func uncommonFromSentences(A string, B string) []string {
	aStrings := strings.Split(A, " ")
	bStrings := strings.Split(B, " ")
	allWords := append(aStrings, bStrings...)

	//重複削除の処理
	wordCntMap := map[string]int{}
	for _, w := range allWords {

		_, ok := wordCntMap[w]
		if !ok {
			wordCntMap[w] = 0
		}
		wordCntMap[w]++
	}

	var res []string
	for s, i := range wordCntMap {
		if i == 1 {
			res = append(res, s)
		}
	}

	return res
}
