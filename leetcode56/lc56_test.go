package leetcode56

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	type Input struct {
		intervals [][]int
		expect    [][]int
	}
	testInput := []Input{
		{
			[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			[][]int{
				{1, 4},
				{4, 5},
			},
			[][]int{
				{1, 5},
			},
		},
	}

	for _, input := range testInput {
		result := Merge(input.intervals)
		assert.Equal(t, input.expect, result)
	}
}
