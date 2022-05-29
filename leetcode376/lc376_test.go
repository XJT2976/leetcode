package leetcode376

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWiggleMaxLength(t *testing.T) {
	type Input struct {
		nums   []int
		expect int
	}
	testInput := []Input{
		{
			[]int{1, 7, 4, 9, 2, 5},
			6,
		},
		{
			[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			2,
		},
	}

	for _, input := range testInput {
		result := WiggleMaxLength(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
