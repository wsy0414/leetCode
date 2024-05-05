package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Given the roots of two binary trees p and q, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
*/
func main() {
	testCases := []struct {
		p      *TreeNode
		q      *TreeNode
		result bool
	}{
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			result: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			result: false,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			result: false,
		},
	}

	for idx, testCase := range testCases {
		if isSameTree(testCase.p, testCase.q) != testCase.result {
			fmt.Printf("seq: %d is fail\n", idx+1)
		}
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}
	// 檢查當下的值和左右邊的值都須相等
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
