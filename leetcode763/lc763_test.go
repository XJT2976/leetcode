package leetcode763

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	type Input struct {
		s      string
		expect []int
	}
	testInput := []Input{
		{
			"ababcbacadefegdehijhklij",
			[]int{9, 7, 8},
		},
		{
			"eccbbbbdec",
			[]int{10},
		},
	}

	for _, input := range testInput {
		result := PartitionLabels(input.s)
		assert.Equal(t, input.expect, result)
	}
}
