package leetcode1005

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestSumAfterKNegations(t *testing.T) {
	type Input struct {
		nums   []int
		k      int
		expect int
	}
	testInput := []Input{
		{
			[]int{4, 2, 3},
			1,
			5,
		},
		{
			[]int{3, -1, 0, 2},
			3,
			6,
		},
		{
			[]int{2, -3, -1, 5, -4},
			2,
			13,
		},
	}

	for _, input := range testInput {
		result := LargestSumAfterKNegations(input.nums, input.k)
		assert.Equal(t, input.expect, result)
	}
}
