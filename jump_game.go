package main

import "fmt"

/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/
func main()  {
  testCases := []struct{
    nums []int
    result bool
  }{
    {
      nums: []int{2, 3, 1, 1, 4},
      result: true,
    }, {
      nums: []int{3, 2, 1, 0, 4},
      result: false,
    },
  }

  for _, testCase := range testCases {
    fmt.Printf("expected result = %v and the actual result = %v\n", testCase.result, canJump(testCase.nums))
  }
}

func canJump(nums []int) bool {
  position := 0

  for {
    if position >= len(nums) {
      break
    }

    if nums[position] == 0 {
      break
    }

    if nums[position] > 0 {
      position += nums[position]
    }
  }

  return position >= len(nums)
}
