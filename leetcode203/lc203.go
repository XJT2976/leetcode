package leetcode203

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveElements(head *ListNode, val int) *ListNode {
	virtualNode := &ListNode{
		0,
		head,
	}
	node := virtualNode

	for node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return virtualNode.Next
}
