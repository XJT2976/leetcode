package leetcode435

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	type Input struct {
		intervals [][]int
		expect    int
	}
	testInput := []Input{
		{
			[][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{1, 3},
			},
			1,
		},
		{
			[][]int{
				{1, 2},
				{1, 2},
				{1, 2},
			},
			2,
		},
		{
			[][]int{
				{1, 2},
				{2, 3},
			},
			0,
		},
	}

	for _, input := range testInput {
		result := EraseOverlapIntervals(input.intervals)
		assert.Equal(t, input.expect, result)
	}
}
