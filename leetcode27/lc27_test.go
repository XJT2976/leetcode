package leetcode27

import (
	"main/utils"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type Input struct {
		nums   []int
		val    int
		result []int
		k      int
	}
	testInput := []Input{
		{
			[]int{3, 2, 2, 3},
			3,
			[]int{2, 2},
			2,
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			[]int{0, 1, 4, 0, 3},
			5,
		},
		{
			[]int{},
			2,
			[]int{},
			0,
		},
	}

	for _, input := range testInput {
		result := RemoveElement(input.nums, input.val)
		if result != input.k {
			t.Errorf("result != input.k, result = %v, input.k = %v", result, input.k)
		} else {
			utils.EasySortInt(input.nums[:result])
			utils.EasySortInt(input.result)
			for i := 0; i < input.k; i++ {
				if input.nums[i] != input.result[i] {
					t.Errorf("result.nums != input.nums, result = %v, input.k = %v",
						input.nums[:result], input.result)
					break
				}
			}
		}
	}
}
