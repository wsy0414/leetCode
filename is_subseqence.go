package main

import (
	"fmt"
	"strings"
)

/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
*/
func main()  {
  testCases := []struct{
    s string
    t string
    result bool
  }{
    {
      s: "abc",
      t: "anbgdc",
      result: true,
    }, {
      s: "axc",
      t: "ahbgdc",
      result: false,
    },
  }

  for _, testCase := range testCases {
    fmt.Printf("expected result = %v and the actual result = %v\n", testCase.result, isSubsquence(testCase.s, testCase.t))
  }
}

func isSubsquence(s string, t string) bool {
  idx := 0
  for _, str := range t {
    if strings.Index(s, string(str)) == idx {
      idx += 1
    } else if strings.Index(s, string(str)) > idx {
      return false
    }
  }

  return idx == len(s)
}
