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

func divideAndMod(num int) (int, int) {
	quotient := num / 10  // 10으로 나눈 몫
	remainder := num % 10 // 10으로 나눈 나머지
	return quotient, remainder
}

func sumTwoLists(l1 *ListNode, l2 *ListNode, pass_val int, result_list *ListNode) *ListNode {

	sum := 0
	if l1 != nil {
		sum += l1.Val
	}
	if l2 != nil {
		sum += l2.Val
	}

	if pass_val > 0 {
		sum += pass_val
		pass_val = 0
	}

	if sum >= 10 {
		quotient, remainder := divideAndMod(sum)
		pass_val = quotient
		sum = remainder
	}

	result_list.Val = sum
	if pass_val == 0 && (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
		result_list.Next = nil
	} else {
		result_list.Next = &ListNode{}
	}

	if (l1 != nil && l1.Next != nil) && (l2 != nil && l2.Next != nil) {
		sumTwoLists(l1.Next, l2.Next, pass_val, result_list.Next)
	} else if (l1 == nil || l1.Next == nil) && (l2 != nil && l2.Next != nil) {
		sumTwoLists(nil, l2.Next, pass_val, result_list.Next)
	} else if (l1 != nil && l1.Next != nil) && (l2 == nil || l2.Next == nil) {
		sumTwoLists(l1.Next, nil, pass_val, result_list.Next)
	} else if pass_val > 0 && (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
		sumTwoLists(nil, nil, pass_val, result_list.Next)
	}

	return result_list

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return sumTwoLists(l1, l2, 0, &ListNode{})
}

// 테스트 함수
func test() {

	// 첫 번째 연결 리스트: 2->4->3 (342)
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 8}}}}

	fmt.Print("l1 : ")
	for current := l1; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

	// 두 번째 연결 리스트: 5->6->4 (465)
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	fmt.Print("l2 : ")
	for current := l2; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

	// 두 수 더하기
	result_sum := addTwoNumbers(l1, l2)

	fmt.Print("결과 : ")
	for current := result_sum; current != nil; current = current.Next {
		fmt.Print(current.Val, " ")
	}
	fmt.Print("\n")

}
