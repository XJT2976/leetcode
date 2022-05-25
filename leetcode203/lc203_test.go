package leetcode203

import (
	"testing"
)

func TestRemoveElements(t *testing.T) {
	type Input struct {
		num    []int
		val    int
		expect []int
	}
	testInput := []Input{
		{
			[]int{1, 2, 6, 3, 4, 5, 6},
			6,
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{},
			1,
			[]int{},
		},
		{
			[]int{7, 7, 7, 7},
			7,
			[]int{},
		},
	}

	for _, input := range testInput {
		inputLink := ConvertArrayToLink(input.num)
		result := ConvertLinkToArray(RemoveElements(inputLink, input.val))

	loop:
		for i := 0; i < len(result); i++ {
			if result[i] != input.expect[i] {
				t.Errorf("result != input.expect, result = %v, input.expect = %v", result, input.expect)
				break loop
			}
		}
	}
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
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}
