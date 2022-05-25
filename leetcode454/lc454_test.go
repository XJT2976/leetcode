package leetcode454

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFourSumCount(t *testing.T) {
	type Input struct {
		nums1  []int
		nums2  []int
		nums3  []int
		nums4  []int
		expect int
	}

	testInput := []Input{
		{
			[]int{1, 2},
			[]int{-2, -1},
			[]int{-1, 2},
			[]int{0, 2},
			2,
		},
		{
			[]int{0},
			[]int{0},
			[]int{0},
			[]int{0},
			1,
		},
	}

	for _, input := range testInput {
		result := FourSumCount(input.nums1, input.nums2, input.nums3, input.nums4)
		assert.Equal(t, input.expect, result)
	}
}
