package codetop

func reverseKGroup(head *ListNode, k int) *ListNode {
	emptyHead := &ListNode{
		Next: head,
	}
	left, right := emptyHead, emptyHead.Next

	reserve := func(left, right *ListNode) (subTail *ListNode) {
		tmpHead := right
		subTail = left.Next
		for left.Next != right {
			tmp := left.Next
			left.Next = tmp.Next
			tmp.Next = tmpHead
			tmpHead = tmp
		}
		left.Next = tmpHead
		return subTail
	}

	for right != nil {
		for i := 0; i < k; i++ {
			if right == nil {
				return emptyHead.Next
			}
			right = right.Next
		}
		left = reserve(left, right)
	}
	return emptyHead.Next
}
