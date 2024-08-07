package leetcode23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	ll := []List{
		{
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
			&ListNode{
				1,
				&ListNode{
					3,
					&ListNode{
						4,
						nil,
					},
				},
			},
			&ListNode{
				2,
				&ListNode{
					6,
					nil,
				},
			},
		},

		{
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
			nil,
		},
		{
			&ListNode{
				1,
				&ListNode{
					4,
					&ListNode{
						10,
						nil,
					},
				},
			},
			nil,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						100,
						nil,
					},
				},
			},
		},
	}

	expects := [][]int{
		{1,1,2,3,4,4,5,6},
		{1,4,5},
		{1,2,3,4,10,100},
	}
	for i := range ll {
		head := MergeKLists(ll[i])
		expect := expects[i]
		for _, j := range expect {
			assert.Equal(t, j, head.Val)
			head = head.Next
		}
	}
}