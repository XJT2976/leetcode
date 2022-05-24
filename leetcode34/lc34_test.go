package leetcode34

import "testing"

func TestSearchRange(t *testing.T) {
	type Input struct {
		nums   []int
		target int
		result []int
	}
	testInput := []Input{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			6,
			[]int{-1, -1},
		},
		{
			[]int{},
			0,
			[]int{-1, -1},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			7,
			[]int{1, 2},
		},
	}

	for _, input := range testInput {
		result := SearchRange(input.nums, input.target)
		if len(result) != len(input.result) || result[0] != input.result[0] ||
			result[1] != input.result[1] {
			t.Errorf("result != input.result, result = %v, input.result = %v", result, input.result)
		}
	}
}
