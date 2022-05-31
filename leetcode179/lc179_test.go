package leetcode179

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	type Input struct {
		nums   []int
		expect string
	}
	testInput := []Input{
		{
			[]int{10, 2},
			"210",
		},
		{
			[]int{3, 30, 34, 5, 9},
			"9534330",
		},
		{
			[]int{1, 3, 4, 5, 123, 12315555},
			"543123155551231",
		},
		{
			[]int{8308, 8308, 830},
			"83088308830",
		},
	}

	for _, input := range testInput {
		result := LargestNumber(input.nums)
		assert.Equal(t, input.expect, result)
	}
}
