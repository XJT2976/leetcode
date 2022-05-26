package leetcode151

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	type Input struct {
		s      string
		expect string
	}
	testInput := []Input{
		{
			"the sky is blue",
			"blue is sky the",
		},
		{
			"  hello world  ",
			"world hello",
		},
		{
			"a good   example",
			"example good a",
		},
		{
			"a",
			"a",
		},
		{
			"",
			"",
		},
	}

	for _, input := range testInput {
		result := ReverseWords(input.s)
		assert.Equal(t, input.expect, result)
	}
}
