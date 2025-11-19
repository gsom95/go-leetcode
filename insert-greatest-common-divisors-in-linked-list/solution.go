package insertgreatestcommondivisorsinlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		next := cur.Next
		a, b := cur.Val, next.Val
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		cur.Next = &ListNode{
			Val:  a,
			Next: next,
		}
		cur = next
	}

	return head
}
