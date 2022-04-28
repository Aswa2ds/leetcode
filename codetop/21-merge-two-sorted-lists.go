package codetop

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	emptyHead := new(ListNode)
	tail := emptyHead
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		list1 = list2
	}
	tail.Next = list1
	return emptyHead.Next
}
