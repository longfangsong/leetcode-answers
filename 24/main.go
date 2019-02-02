package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oldHead := head
	head = head.Next
	oldHead.Next = head.Next
	head.Next = oldHead
	current := head.Next
	currentNext := current.Next
	for currentNext != nil && currentNext.Next != nil {
		currentNextNext := currentNext.Next
		current.Next = currentNextNext
		currentNext.Next = currentNextNext.Next
		currentNextNext.Next = currentNext
		current = current.Next.Next
		currentNext = current.Next
	}
	return head
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	fmt.Println(l1.Val, l1.Next.Val, l1.Next.Next.Val)
	// swapPairs(l1)
}
