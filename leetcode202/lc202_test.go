package leetcode202

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsHappy(t *testing.T) {
	type Input struct {
		n      int
		expect bool
	}

	testInput := []Input{
		{
			19,
			true,
		},
		{
			2,
			false,
		},
		{
			1,
			true,
		},
	}

	for _, input := range testInput {
		result := IsHappy(input.n)
		assert.Equal(t, input.expect, result)
	}
}
