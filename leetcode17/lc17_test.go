package leetcode17

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	type Input struct {
		digits string
		expect []string
	}
	testInput := []Input{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"",
			[]string{},
		},
	}

	for _, input := range testInput {
		result := LetterCombinations(input.digits)
		assert.Equal(t, input.expect, result)
	}
}
