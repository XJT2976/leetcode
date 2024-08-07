package leetcode86


type ListNode struct {
	Val int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	var now, lessHead, lessTail, greatHead, greatTail *ListNode
	now = head
	for ;now != nil; {
		if now.Val < x {
			if lessHead == nil {
				lessHead = now
				lessTail = now
			} else {
				lessTail.Next = now
				lessTail = now
			}
		} else{
			if greatHead == nil {
				greatHead = now
				greatTail = now
			} else {
				greatTail.Next = now
				greatTail = now
			}
		}

		now = now.Next
	}

	if lessTail != nil {
		lessTail.Next = greatHead
	}

	if greatTail != nil {
		greatTail.Next = nil
	}

	if lessHead != nil {
		return lessHead
	} else {
		return greatHead
	}
}