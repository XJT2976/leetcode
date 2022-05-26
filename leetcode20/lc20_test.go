package leetcode20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	type Input struct {
		s      string
		expect bool
	}
	testInput := []Input{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"()[{}",
			false,
		},
		{
			"(()",
			false,
		},
	}

	for _, input := range testInput {
		result := IsValid(input.s)
		assert.Equal(t, input.expect, result)
	}
}
