package palindromelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var elements []int
	for head != nil {
		elements = append(elements, head.Val)
		head = head.Next
	}

	size := len(elements)
	for i := 0; i < size/2; i++ {
		if elements[i] != elements[size-1-i] {
			return false
		}
	}

	return true
}
