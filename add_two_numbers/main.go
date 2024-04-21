package main

import "fmt"

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	addTwoNumbers(&l1, &l2).toString()

	l1 = ListNode{
		Val: 0,
	}

	l2 = ListNode{
		Val: 0,
	}

	addTwoNumbers(&l1, &l2).toString()

	l1 = ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}

	l2 = ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}

	addTwoNumbers(&l1, &l2).toString()
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNode *ListNode) toString() string {
	result := ""
	for listNode != nil {
		fmt.Printf("%d", listNode.Val)
		listNode = listNode.Next
	}
	fmt.Printf("\n")
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	// 透過copy by pointer的方式 讓修改temp等於修改result
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val = tmp.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val = tmp.Val + l2.Val
			l2 = l2.Next
		}
		// 大於9需要進位，所以放1到Next讓下一個使用
		if tmp.Val > 9 {
			tmp.Next = &ListNode{
				Val: 1,
			}
			tmp.Val = tmp.Val - 10
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}

	return result
}
