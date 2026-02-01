package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nodeSum(l1, l2, 0)
}

func nodeSum(l1 *ListNode, l2 *ListNode, x int) *ListNode {
	if l1 == nil && l2 == nil && x == 0 {
		return nil
	}
	var (
		val1  int
		next1 *ListNode
		val2  int
		next2 *ListNode
		nextX int
	)
	if l1 != nil {
		val1 = l1.Val
		next1 = l1.Next
	}
	if l2 != nil {
		val2 = l2.Val
		next2 = l2.Next
	}
	val := val1 + val2 + x
	if val >= 10 {
		nextX = 1
		val -= 10
	}
	nextNode := nodeSum(next1, next2, nextX)
	return &ListNode{Val: val, Next: nextNode}
}
