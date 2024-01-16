package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curPtr := head.Next
	for curPtr != nil {
		if curPtr.Val == prev.Val {
			prev.Next = curPtr.Next
		} else {
			prev = curPtr
		}
		curPtr = curPtr.Next
	}

	return head
}
