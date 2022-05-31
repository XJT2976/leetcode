package leetcode55

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	type Input struct {
		prices []int
		expect bool
	}
	testInput := []Input{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{1, 1, 1, 1, 1},
			true,
		},
	}

	for _, input := range testInput {
		result := CanJump(input.prices)
		assert.Equal(t, input.expect, result)
	}
}
