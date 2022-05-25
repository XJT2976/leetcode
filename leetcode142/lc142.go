package leetcode142

import "main/utils"

func DetectCycle(head *utils.ListNode) *utils.ListNode {
	slow := head
	fast := head

	for {
		if fast == nil || fast.Next == nil {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}

	return nil
}
