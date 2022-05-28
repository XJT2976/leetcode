package leetcode131

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	type Input struct {
		s      string
		expect [][]string
	}
	testInput := []Input{
		{
			"cbbbcc",
			[][]string{
				{"c", "b", "b", "b", "c", "c"},
				{"c", "b", "b", "b", "cc"},
				{"c", "b", "bb", "c", "c"},
				{"c", "b", "bb", "cc"},
				{"c", "bb", "b", "c", "c"},
				{"c", "bb", "b", "cc"},
				{"c", "bbb", "c", "c"},
				{"c", "bbb", "cc"},
				{"cbbbc", "c"},
			},
		},
		{
			"aab",
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			"a",
			[][]string{
				{"a"},
			},
		},
		{
			"abacdc",
			[][]string{
				{"a", "b", "a", "c", "d", "c"},
				{"a", "b", "a", "cdc"},
				{"aba", "c", "d", "c"},
				{"aba", "cdc"},
			},
		},
	}

	for _, input := range testInput {
		result := Partition(input.s)
		assert.Equal(t, input.expect, result)
	}
}
