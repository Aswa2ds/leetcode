package codetop

func reverseList(head *ListNode) *ListNode {
	emptyHead := &ListNode{}
	emptyHead.Next = head
	var newHead *ListNode = nil
	for emptyHead.Next != nil {
		tmp := emptyHead.Next
		emptyHead.Next = tmp.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}
