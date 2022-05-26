package leetcode150

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	type Input struct {
		s      []string
		expect int
	}
	testInput := []Input{
		{
			[]string{"2", "1", "+", "3", "*"},
			9,
		},
		{
			[]string{"4", "13", "5", "/", "+"},
			6,
		},
		{
			[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			22,
		},
	}

	for _, input := range testInput {
		result := EvalRPN(input.s)
		assert.Equal(t, input.expect, result)
	}
}
