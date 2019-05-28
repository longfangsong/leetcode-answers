package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseOneGroup(beforeBegin *ListNode, beforeEnd *ListNode) {
	newTail := beforeBegin.Next
	end := beforeEnd.Next
	for newTail.Next != end {
		toMoveFront := newTail.Next
		newTail.Next = toMoveFront.Next
		toMoveFront.Next = beforeBegin.Next
		beforeBegin.Next = toMoveFront
	}
}

func advance(node *ListNode, k int) *ListNode {
	for i := 0; i < k; i++ {
		node = node.Next
		if node == nil {
			return nil
		}
	}
	return node
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	beforeHead := &ListNode{0, head}
    beforeBegin := beforeHead
    for {
        beforeEnd := advance(beforeBegin,k)
        if beforeEnd == nil {
            break;
        }
        reverseOneGroup(beforeBegin,beforeEnd)
        beforeBegin = advance(beforeEnd,k-1)
    }
    return beforeHead.Next
}

func main() {
	l := &ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	println(advance(l, 2).Val)
}
