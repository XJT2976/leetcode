package leetcode53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	type Input struct {
		nums   []int
		expect int
	}
	testInput := []Input{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}

	for _, input := range testInput {
		result := MaxSubArray(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
