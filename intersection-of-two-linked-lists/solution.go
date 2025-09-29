package intersectionoftwolinkedlists


type ListNode struct {
    Val int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    elems := map[*ListNode]struct{}{}

    a, b := headA, headB
    for a != nil || b != nil {
        if a != nil {
            if _, seen := elems[a]; seen {
                return a
            }
            elems[a] = struct{}{}
            a = a.Next
        }
        if b != nil {
            if _, seen := elems[b]; seen {
                return b
            }
            elems[b] = struct{}{}
            b = b.Next
        }
    }

    return nil
}