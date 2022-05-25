package leetcode242

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type Input struct {
		s      string
		t      string
		expect bool
	}

	testInput := []Input{
		{
			"anagram",
			"nagaram",
			true,
		},
		{
			"rat",
			"car",
			false,
		},
		{
			"anagram",
			"nagnram",
			false,
		},
	}

	for _, input := range testInput {
		result := IsAnagram(input.s, input.t)
		assert.Equal(t, input.expect, result)
	}
}
