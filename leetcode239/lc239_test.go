package leetcode239

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	type Input struct {
		nums   []int
		k      int
		expect []int
	}
	testInput := []Input{
		{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, input := range testInput {
		result := MaxSlidingWindow(input.nums, input.k)
		assert.Equal(t, input.expect, result)
	}
}
