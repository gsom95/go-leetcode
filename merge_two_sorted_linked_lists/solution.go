package mergetwosortedlinkedlists

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{}
	curList := head

	for {
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				curList.Val = list1.Val
				list1 = list1.Next
			} else {
				curList.Val = list2.Val
				list2 = list2.Next
			}
		} else if list2 == nil {
			curList.Val = list1.Val
			list1 = list1.Next
		} else if list1 == nil {
			curList.Val = list2.Val
			list2 = list2.Next
		}

		if list1 == nil && list2 == nil {
			break
		}
		curList.Next = &ListNode{}
		curList = curList.Next

	}

	return head
}
