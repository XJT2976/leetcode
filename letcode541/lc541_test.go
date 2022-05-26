package letcode541

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseStr(t *testing.T) {
	type Input struct {
		s      string
		k      int
		expect string
	}
	testInput := []Input{
		{
			"abcdefg",
			2,
			"bacdfeg",
		},
		{
			"abcd",
			2,
			"bacd",
		},
		{
			"ab",
			4,
			"ba",
		},
		{
			"abc",
			4,
			"cba",
		},
		{
			"abcdefghijk",
			4,
			"dcbaefghkji",
		},
	}

	for _, input := range testInput {
		result := ReverseStr(input.s, input.k)
		assert.Equal(t, input.expect, result)
	}
}
