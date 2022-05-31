package leetcode122

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	type Input struct {
		prices []int
		expect int
	}
	testInput := []Input{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	for _, input := range testInput {
		result := MaxProfit(input.prices)
		assert.Equal(t, input.expect, result)
	}
}
