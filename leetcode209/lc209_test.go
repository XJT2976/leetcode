package leetcode209

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	type Input struct {
		nums   []int
		target int
		expect int
	}
	testInput := []Input{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			11,
			0,
		},
		{
			[]int{},
			7,
			0,
		},
		{
			[]int{7},
			7,
			1,
		},
		{
			[]int{1, 7},
			7,
			1,
		},
		{
			[]int{10, 2, 3, 100},
			7,
			1,
		},
	}

	for _, input := range testInput {
		result := MinSubArrayLen(input.target, input.nums)
		if result != input.expect {
			t.Errorf("result != input.expect, result = %v, input.expect = %v", result, input.expect)
			break
		}
	}
}
