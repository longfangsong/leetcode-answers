package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	cutAt := head
	nullDetector := head
	for i := 0; i < k; i++ {
		if nullDetector.Next == nil {
			nullDetector = head
		} else {
			nullDetector = nullDetector.Next
		}
	}
	if nullDetector == cutAt {
		return head
	}
	for nullDetector.Next != nil {
		nullDetector = nullDetector.Next
		cutAt = cutAt.Next
	}
	newHead := cutAt.Next
	cutAt.Next = nil
	nullDetector.Next = head
	return newHead
}
