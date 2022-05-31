package leetcode452

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMinArrowShots(t *testing.T) {
	type Input struct {
		points [][]int
		expect int
	}
	testInput := []Input{
		{
			[][]int{
				{10, 16},
				{2, 8},
				{1, 6},
				{7, 12},
			},
			2,
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			},
			4,
		},
		{
			[][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
			},
			2,
		},
	}

	for _, input := range testInput {
		result := FindMinArrowShots(input.points)
		assert.Equal(t, input.expect, result)
	}
}
