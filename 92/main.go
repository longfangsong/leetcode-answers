package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		newNode := new(ListNode)
		newNode.Next = head
		return reverseBetween(newNode, 2, n+1).Next
	}
	cutAt := head
	for i := 1; i < m-1; i++ {
		cutAt = cutAt.Next
	}
	reversed := cutAt.Next
	originNext := cutAt.Next
	toReverse := reversed.Next
	for i := m; i < n; i++ {
		newToReverse := toReverse.Next
		toReverse.Next = reversed
		reversed = toReverse
		toReverse = newToReverse
	}
	cutAt.Next = reversed
	originNext.Next = toReverse
	return head
}

func main() {
	newNode := new(ListNode)
	newNode.Val = 1
	newNode.Next = new(ListNode)
	newNode.Next.Val = 2
	newNode.Next.Next = nil
	newNode = reverseBetween(newNode, 1, 2)
	println(newNode.Val)
	println(newNode.Next.Val)
}
