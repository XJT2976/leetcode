package leetcode86

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	l := []*ListNode{
		{
			1,
			&ListNode{
				4,
				&ListNode{
					3,
					&ListNode{
						2,
						&ListNode{
							5,
							&ListNode{
								2,
								nil,
							},
						},
					},
				},
			},
		},
		{
			2,
			&ListNode{
				1,
				nil,
			},
		},
		nil,
		{10, nil},
		{100, nil},
		{20, &ListNode{10, nil}},
		{150, &ListNode{200, nil}},
	}
	expects := [][]int{
		{1, 2, 2, 4, 3, 5},
		{1, 2},
		{},
		{10},
		{100},
		{20, 10},
		{150, 200},
	}
	xl := []int{3, 2, 5, 50, 50, 50, 50}
	for i := range l {
		head := Partition(l[i], xl[i])
		for j := range expects[i] {
			assert.Equal(t, expects[i][j], head.Val)
			head = head.Next
		}
	}
}
