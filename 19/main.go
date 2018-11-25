package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{0, head}
	cursor1 := &dummy
	cursor2 := &dummy
	for i := 0; i <= n; i++ {
		cursor2 = cursor2.Next
	}
	for cursor2 != nil {
		cursor1 = cursor1.Next
		cursor2 = cursor2.Next
	}
	cursor1.Next = cursor1.Next.Next
	return dummy.Next
}

func main() {
	l := ListNode{1, nil}
	removeNthFromEnd(&l, 1)
}
