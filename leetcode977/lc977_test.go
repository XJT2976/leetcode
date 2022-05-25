package leetcode977

import "testing"

func TestSortedSquares(t *testing.T) {
	type Input struct {
		nums   []int
		expect []int
	}
	testInput := []Input{
		{
			[]int{-4, -1, 0, 3, 10},
			[]int{0, 1, 9, 16, 100},
		},
		{
			[]int{-7, -3, 2, 3, 11},
			[]int{4, 9, 9, 49, 121},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{0},
			[]int{0},
		},
		{
			[]int{-1},
			[]int{1},
		},
		{
			[]int{2},
			[]int{4},
		},
		{
			[]int{0, 3, 10},
			[]int{0, 9, 100},
		},
		{
			[]int{-3, -2, 0},
			[]int{0, 4, 9},
		},
		{
			[]int{-3, -2},
			[]int{4, 9},
		},
		{
			[]int{3, 10},
			[]int{9, 100},
		},
	}

	for _, input := range testInput {
		result := SortedSquares(input.nums)

		for i := range result {
			if result[i] != input.expect[i] {
				t.Errorf("result != input.expect, result = %v, input.expect = %v", result, input.expect)
				break
			}
		}
	}
}
