package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 2, 3}
	sum := removeElement(nums, 3)

	fmt.Println(nums)
	fmt.Println(sum)
}

// 用類似排序方式將不等於val得值放到陣列前面，因為測試時只排序不符合val的數量長度
func removeElement(nums []int, val int) int {
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
