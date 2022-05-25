package leetcode206

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	cur := head
	prev := (*ListNode)(nil)
	next := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func ReverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := ReverseListRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return ret
}
