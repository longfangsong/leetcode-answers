func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
        if slow == fast {
            break
        }
	}
	result := head
	for result != slow {
		slow = slow.Next
		result = result.Next
	}
	return result
}