package removelinkedlistelements

type ListNode struct {
    Val int
    Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
    newHead := head
    for newHead != nil {
        if newHead.Val != val {
            break
        }
        newHead = newHead.Next
    }

    var (
        cur = newHead
        curNext *ListNode
    )
    for cur != nil {
        curNext = cur.Next
        if curNext == nil {
            break
        }
        if curNext.Val == val {
            cur.Next = curNext.Next
            continue
        }

        cur = cur.Next
    }

    return newHead
}