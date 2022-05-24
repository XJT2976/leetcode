package leetcode35

import "testing"

func TestSearchInsert(t *testing.T) {
	type Input struct {
		nums   []int
		target int
		result int
	}
	testInput := []Input{
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			[]int{1},
			7,
			1,
		},
		{
			[]int{11},
			7,
			0,
		},
	}

	for _, input := range testInput {
		result := SearchInsert(input.nums, input.target)
		if result != input.result {
			t.Errorf("result != input.result, result = %v, input.result = %v", result, input.result)
		}
	}
}
