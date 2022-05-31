package leetcode135

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandy(t *testing.T) {
	type Input struct {
		ratings []int
		expect  int
	}
	testInput := []Input{
		{
			// 2 1 2
			[]int{1, 0, 2},
			5,
		},
		{
			[]int{1, 2, 2},
			4,
		},
		{
			[]int{100},
			1,
		},
		{
			// 1 2 1 2 1
			[]int{1, 3, 2, 2, 1},
			7,
		},
		{
			// 1 2 3 1 1
			// 1 2 3 2 1
			[]int{1, 2, 3, 1, 0},
			9,
		},
	}

	for _, input := range testInput {
		result := Candy(input.ratings)
		assert.Equal(t, input.expect, result)
	}
}
