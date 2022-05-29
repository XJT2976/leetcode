package leetcode90

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	type Input struct {
		nums   []int
		expect [][]int
	}
	testInput := []Input{
		{
			[]int{1, 2, 2},
			[][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 2},
				{2},
				{2, 2},
			},
		},
		{
			[]int{0},
			[][]int{
				{},
				{0},
			},
		},
	}

	for _, input := range testInput {
		result := SubsetsWithDup(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
