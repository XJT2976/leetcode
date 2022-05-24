package leetcode704

import "testing"

func TestSearch(t *testing.T) {
	type Input struct {
		nums   []int
		target int
		result int
	}
	testInput := []Input{
		{
			[]int{-1, 0, 3, 5, 9, 12},
			9,
			4,
		},
		{
			[]int{-1, 0, 3, 5, 9, 12},
			2,
			-1,
		},
		{
			[]int{},
			2,
			-1,
		},
		{
			[]int{1},
			2,
			-1,
		},
		{
			[]int{1},
			1,
			0,
		},
	}

	for _, input := range testInput {
		result := Search(input.nums, input.target)
		if result != input.result {
			t.Errorf("result != input.result, result = %v, input.result = %v", result, input.result)
		}
	}
}
