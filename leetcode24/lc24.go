package leetcode24

import "main/utils"

func SwapPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nNext := SwapPairs(head.Next.Next)
	head.Next.Next = head
	ret := head.Next
	head.Next = nNext
	return ret
}
