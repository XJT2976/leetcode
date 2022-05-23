package leetcode1

import "testing"

func TestTwoSum(t *testing.T) {
	type Input struct {
		nums   []int
		target int
		result []int
	}
	testInput := []Input{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
	}

	for _, input := range testInput {
		result := TwoSum(input.nums, input.target)
		if !(len(result) == 2 && ((result[0] == input.result[0] &&
			result[1] == input.result[1]) || (result[0] == input.result[1] &&
			result[1] == input.result[0]))) {
			t.Errorf("result != input.result, result = %v, input.result = %v", result, input.result)
		}
	}
}
