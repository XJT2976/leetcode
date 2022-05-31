package leetcode62

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	type Input struct {
		m      int
		n      int
		expect int
	}
	testInput := []Input{
		{
			3,
			7,
			28,
		},
		{
			3,
			2,
			3,
		},
	}

	for _, input := range testInput {
		result := UniquePaths(input.m, input.n)
		assert.Equal(t, input.expect, result)
	}
}
