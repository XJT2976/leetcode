package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func ConvertArrayToLink(nums []int) *ListNode {
	head := (*ListNode)(nil)
	prev := head
	for _, i := range nums {
		if head == nil {
			head = &ListNode{
				i,
				nil,
			}
			prev = head
		} else {
			node := &ListNode{
				i,
				nil,
			}
			prev.Next = node
			prev = node
		}
	}
	return head
}

func ConvertLinkToArray(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}

func EasySortInt(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
