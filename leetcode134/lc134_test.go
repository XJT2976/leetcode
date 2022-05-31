package leetcode134

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	type Input struct {
		gas    []int
		cost   []int
		expect int
	}
	testInput := []Input{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
	}

	for _, input := range testInput {
		result := CanCompleteCircuit(input.gas, input.cost)
		assert.Equal(t, input.expect, result)
	}
}
