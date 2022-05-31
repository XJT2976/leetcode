package leetcode343

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegerBreak(t *testing.T) {
	type Input struct {
		n      int
		expect int
	}
	testInput := []Input{
		{
			2,
			1,
		},
		{
			10,
			36,
		},
		{
			// 3 6
			9,
			27,
		},
		{
			// 3 6
			14,
			162,
		},
	}

	for _, input := range testInput {
		result := IntegerBreak(input.n)
		assert.Equal(t, input.expect, result)
	}
}
