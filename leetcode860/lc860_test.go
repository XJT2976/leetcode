package leetcode860

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLemonadeChange(t *testing.T) {
	type Input struct {
		bills  []int
		expect bool
	}
	testInput := []Input{
		{
			[]int{5, 5, 5, 10, 20},
			true,
		},
		{
			[]int{5, 5, 10, 10, 20},
			false,
		},
	}

	for _, input := range testInput {
		result := LemonadeChange(input.bills)
		assert.Equal(t, input.expect, result)
	}
}
