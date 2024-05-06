package main

import "fmt"

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	// 預設0在index=0的位置
	zeroIdx := 0
	for i := 0; i < len(nums); i++ {
		// 遇到非0數字後和0互換，然後調整zeroIdx的數字
		if nums[i] != 0 {
			nums[i], nums[zeroIdx] = nums[zeroIdx], nums[i]
			zeroIdx++
		}
		fmt.Println(nums)
	}
}
