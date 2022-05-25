package leetcode142

import (
	"github.com/stretchr/testify/assert"
	"main/utils"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	type Input struct {
		head   *utils.ListNode
		expect *utils.ListNode
	}
	nodes := make([]*utils.ListNode, 10)
	for i := 0; i < 10; i++ {
		nodes[i] = &utils.ListNode{}
	}

	nodes[0].Next = nodes[1]
	nodes[1].Next = nodes[2]
	nodes[2].Next = nodes[3]
	nodes[3].Next = nodes[1]

	nodes[4].Next = nodes[5]
	nodes[5].Next = nodes[4]

	nodes[7].Next = nodes[8]
	nodes[8].Next = nodes[9]

	testInput := []Input{
		{
			nodes[0],
			nodes[1],
		},
		{
			nodes[4],
			nodes[4],
		},
		{
			nodes[6],
			nil,
		},
		{
			nodes[7],
			nil,
		},
	}

	for _, input := range testInput {
		result := DetectCycle(input.head)
		assert.Equal(t, input.expect, result)
	}
}
