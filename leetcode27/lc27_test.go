package leetcode27

import "testing"

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
			easySortInt(input.nums[:result])
			easySortInt(input.result)
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

func easySortInt(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
