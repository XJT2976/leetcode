package leetcode738

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMonotoneIncreasingDigits(t *testing.T) {
	type Input struct {
		n      int
		expect int
	}
	testInput := []Input{
		{
			10,
			9,
		},
		{
			1234,
			1234,
		},
		{
			332,
			299,
		},
	}

	for _, input := range testInput {
		result := MonotoneIncreasingDigits(input.n)
		assert.Equal(t, input.expect, result)
	}
}
