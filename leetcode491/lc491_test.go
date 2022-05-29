package leetcode491

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	type Input struct {
		nums   []int
		expect [][]int
	}
	testInput := []Input{
		{
			[]int{4, 6, 7, 7},
			[][]int{
				{4, 6},
				{4, 6, 7},
				{4, 6, 7, 7},
				{4, 7},
				{4, 7, 7},
				{6, 7},
				{6, 7, 7},
				{7, 7},
			},
		},
		{
			[]int{4, 4, 3, 2, 1},
			[][]int{
				{4, 4},
			},
		},
	}

	for _, input := range testInput {
		result := FindSubsequences(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
