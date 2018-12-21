package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cursor1 := l1
	cursor2 := l2
	var resultFirst *ListNode
	var resultLast *ListNode
	if l1.Val < l2.Val {
		resultFirst = cursor1
		resultLast = cursor1
		cursor1 = cursor1.Next
	} else {
		resultFirst = cursor2
		resultLast = cursor2
		cursor2 = cursor2.Next
	}
	for cursor1 != nil && cursor2 != nil {
		if cursor1.Val < cursor2.Val {
			resultLast.Next = cursor1
			resultLast = resultLast.Next
			cursor1 = cursor1.Next
		} else {
			resultLast.Next = cursor2
			resultLast = resultLast.Next
			cursor2 = cursor2.Next
		}
	}
	if cursor1 != nil {
		resultLast.Next = cursor1
	} else {
		resultLast.Next = cursor2
	}
	return resultFirst
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	mergeTwoLists(l1, l2)
}
