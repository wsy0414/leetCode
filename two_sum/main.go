package main

import "fmt"

func main() {
	fmt.Printf("output: %v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("output: %v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("output: %v\n", twoSum([]int{3, 3}, 6))
	fmt.Printf("output: %v\n", twoSumWithMap([]int{2, 7, 11, 15}, 9))
	fmt.Printf("output: %v\n", twoSumWithMap([]int{3, 2, 4}, 6))
	fmt.Printf("output: %v\n", twoSumWithMap([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}

	return []int{}
}

func twoSumWithMap(nums []int, target int) []int {
	var s = make(map[int]int)

	// 如果key取不到值 ok 會是 false
	for idx, num := range nums {
		if pos, ok := s[target-num]; ok {
			return []int{pos, idx}
		}
		s[num] = idx
	}

	return []int{}
}
