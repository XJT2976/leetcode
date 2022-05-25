package leetcode19

import (
	"github.com/stretchr/testify/assert"
	"main/utils"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	type Input struct {
		num    []int
		index  int
		expect []int
	}
	testInput := []Input{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{1, 2, 3, 5},
		},
		{
			[]int{1},
			1,
			[]int{},
		},
		{
			[]int{},
			1,
			[]int{},
		},
		{
			[]int{1, 2},
			1,
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			1,
			[]int{1, 2, 3, 4},
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
			[]int{2, 3, 4, 5},
		},
	}

	for _, input := range testInput {
		inputLink := utils.ConvertArrayToLink(input.num)
		result := utils.ConvertLinkToArray(removeNthFromEnd(inputLink, input.index))
		assert.Equal(t, input.expect, result)
	}
}
