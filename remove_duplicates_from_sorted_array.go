package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.
Custom Judge:

The judge will test your solution with the following code:

int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;

	for (int i = 0; i < k; i++) {
	    assert nums[i] == expectedNums[i];
	}

If all assertions pass, then your solution will be accepted.
*/
func main() {
	testCases := []struct {
		nums         []int
		expectedNums []int
		result       int
	}{
		{
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
			result:       3,
		},
		{
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
			result:       5,
		},
	}

	for _, testCase := range testCases {
		result := removeDuplicates(testCase.nums)
		if testCase.result != result {
			fmt.Printf("result: %d is not equal to expect: %d\n", result, testCase.result)
			continue
		}
		if result != len(testCase.expectedNums) {
			fmt.Printf("result: %d is not equal length of expectedNums: %d\n", result, len(testCase.expectedNums))
			continue
		}
		fmt.Println(testCase.nums)
		sort.Ints(testCase.nums[0:result])
		for i := 0; i < result; i++ {
			if testCase.nums[i] != testCase.expectedNums[i] {
				fmt.Printf("nums[%d]: %d is not equalt to expectedNums[%d]: %d\n", i, testCase.nums[i], i, testCase.expectedNums[i])
				continue
			}
		}
	}
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	i := 0
	for v, val := range nums {
		if _, ok := m[val]; !ok {
			nums[i] = val
			i++
		}
		m[val] = v
	}

	return i
}
