package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{1}, 1, []int{}, 0)
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	fmt.Printf("nums1 = %v, m = %v, nums2 = %v, n = %v after merge result = %v\n", nums1, m, nums2, n, nums1)
}
