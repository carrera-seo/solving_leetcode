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

func reverseList(ll *ListNode) []int {
	// list 를 뒤집어서 배열에 저장한다.
	var arr []int

	for current := ll; current != nil; current = current.Next {
		arr = append([]int{current.Val}, arr...)
	}

	return arr
}

func divideAndMod(num int) (int, int) {
	quotient := num / 10  // 10으로 나눈 몫
	remainder := num % 10 // 10으로 나눈 나머지
	return quotient, remainder
}

func listToListNode(sum_list []int, result_list *ListNode) *ListNode {

	//fmt.Printf("quo : %d , rem : %d \n", quotient, remainder)

	if len(sum_list) == 1 {
		result_list.Val = sum_list[0]
		result_list.Next = nil
	} else {
		first := sum_list[0]
		sum_list = sum_list[1:]
		result_list.Val = first
		result_list.Next = &ListNode{Val: 0}
		listToListNode(sum_list, result_list.Next)
	}
	return result_list
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	l1_result := reverseList(l1)
	l2_result := reverseList(l2)

	if len(l1_result) > len(l2_result) {
		need_count := len(l1_result) - len(l2_result)
		for i := 0; i < need_count; i++ {
			l2_result = append([]int{0}, l2_result...)
		}
	} else if len(l1_result) < len(l2_result) {
		need_count := len(l2_result) - len(l1_result)
		for i := 0; i < need_count; i++ {
			l1_result = append([]int{0}, l1_result...)
		}
	}

	//two_sum := l1_result + l2_result
	//fmt.Print("l1 : ", l1_result, "\n")
	//fmt.Print("l2 : ", l2_result, "\n")

	var sum_list []int

	next_sum := 0

	for i := len(l1_result) - 1; i >= 0; i-- {

		sum := l1_result[i] + l2_result[i]

		if next_sum > 0 {
			sum += next_sum
			next_sum = 0
		}

		if sum >= 10 {
			quotient, remainder := divideAndMod(sum)
			next_sum = quotient
			sum = remainder
		}

		sum_list = append(sum_list, sum)

	}

	if next_sum > 0 {
		sum_list = append(sum_list, next_sum)
	}

	//fmt.Print("sum_list : ", sum_list, "\n")

	result := listToListNode(sum_list, &ListNode{})

	return result

}

// 테스트 함수
func main() {

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
