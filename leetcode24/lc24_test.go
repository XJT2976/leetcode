package leetcode24

import (
	"github.com/stretchr/testify/assert"
	"main/utils"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	type Input struct {
		num    []int
		expect []int
	}
	testInput := []Input{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 1, 4, 3},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 1, 4, 3, 5},
		},
	}

	for _, input := range testInput {
		inputLink := utils.ConvertArrayToLink(input.num)
		result := utils.ConvertLinkToArray(SwapPairs(inputLink))
		assert.Equal(t, input.expect, result)
	}
}
