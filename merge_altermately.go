package main

import (
	"fmt"
	"strings"
)

/*
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.
*/
func main() {
	testCases := []struct {
		word1  string
		word2  string
		expect string
	}{
		{
			word1:  "abc",
			word2:  "pqr",
			expect: "apbqcr",
		}, {
			word1:  "ab",
			word2:  "pqrs",
			expect: "apbqrs",
		}, {
			word1:  "abcd",
			word2:  "pq",
			expect: "apbqcd",
		},
	}

	for _, testCase := range testCases {
		result := mergeAlternately(testCase.word1, testCase.word2)
		if result != testCase.expect {
			fmt.Printf("result: %s is not equal to expect: %s", result, testCase.expect)
		}
	}
}

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	word1List := strings.Split(word1, "")
	word2List := strings.Split(word2, "")
	i := 0
	// 先放word1在放word2
	for j := 0; j < len(word1List); j++ {
		builder.WriteString(word1List[j])
		// 如果word1還有word2放完就不放了
		if i < len(word2List) {
			builder.WriteString(word2List[i])
			i++
		}
	}
	// 如果word1放完word2還有就把剩下的放完
	if i < len(word2List) {
		for j := i; j < len(word2List); j++ {
			builder.WriteString(word2List[j])
		}
	}
	return builder.String()
}
