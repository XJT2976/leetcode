package leetcode347

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	type Input struct {
		nums   []int
		k      int
		expect []int
	}
	testInput := []Input{
		{
			[]int{1, 1, 1, 2, 2, 3},
			2,
			[]int{1, 2},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
		{
			[]int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4},
			2,
			[]int{4, 1},
		},
	}

	for _, input := range testInput {
		result := TopKFrequent(input.nums, input.k)
		sort.Ints(result)
		sort.Ints(input.expect)
		assert.Equal(t, input.expect, result)
	}
}
