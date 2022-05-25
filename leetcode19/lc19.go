package leetcode19

import "main/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	length := 0
	node := head
	for node != nil {
		node = node.Next
		length++
	}

	if n > length {
		return nil
	}
	if n == length && head != nil {
		return head.Next
	}

	node = head
	for index := 0; index < length-n-1; index++ {
		node = node.Next
	}

	if node.Next != nil {
		node.Next = node.Next.Next
	}

	return head
}
