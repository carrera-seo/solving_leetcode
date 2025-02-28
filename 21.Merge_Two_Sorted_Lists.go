package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	var result_node ListNode

	current := &result_node

	for l1, l2 := list1, list2; true; {

		if l1 == nil && l2 != nil {
			current.Val = l2.Val
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			current.Val = l1.Val
			l1 = l1.Next
		} else if l1.Val < l2.Val {
			current.Val = l1.Val
			l1 = l1.Next
		} else {
			current.Val = l2.Val
			l2 = l2.Next
		}

		if l1 == nil && l2 == nil {
			break
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}

	}

	return &result_node

}

func main() {

	// 첫 번째 연결 리스트: 2->3->5->8
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8}}}}

	fmt.Print("l1 : ")
	for current := l1; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

	// 두 번째 연결 리스트: 4->5->6 (
	l2 := &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}

	fmt.Print("l2 : ")
	for current := l2; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

	// 두 리스트 머지
	result_merge := mergeTwoLists(l1, l2)

	fmt.Print("결과 : ")
	for current := result_merge; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

}
