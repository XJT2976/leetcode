package leetcode45

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJump(t *testing.T) {
	type Input struct {
		prices []int
		expect int
	}
	testInput := []Input{
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
		},
	}

	for _, input := range testInput {
		result := Jump(input.prices)
		assert.Equal(t, input.expect, result)
	}
}
