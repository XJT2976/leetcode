package leetcode455

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	type Input struct {
		g      []int
		s      []int
		expect int
	}
	testInput := []Input{
		{
			[]int{1, 2, 3},
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2},
			[]int{1, 2, 3},
			2,
		},
	}

	for _, input := range testInput {
		result := FindContentChildren(input.g, input.s)
		assert.Equal(t, input.expect, result)
	}
}
