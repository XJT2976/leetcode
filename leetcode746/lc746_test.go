package leetcode746

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	type Input struct {
		cost   []int
		expect int
	}
	testInput := []Input{
		{
			[]int{10, 15, 20},
			15,
		},
		{
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
	}

	for _, input := range testInput {
		result := MinCostClimbingStairs(input.cost)
		assert.Equal(t, input.expect, result)
	}
}
