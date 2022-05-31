package leetcode406

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	type Input struct {
		people [][]int
		expect [][]int
	}
	testInput := []Input{
		{
			[][]int{
				{7, 0},
				{4, 4},
				{7, 1},
				{5, 0},
				{6, 1},
				{5, 2},
			},
			[][]int{
				{5, 0},
				{7, 0},
				{5, 2},
				{6, 1},
				{4, 4},
				{7, 1},
			},
		},
	}

	for _, input := range testInput {
		result := ReconstructQueue(input.people)
		assert.Equal(t, input.expect, result)
	}
}
