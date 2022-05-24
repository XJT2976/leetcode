package leetcode367

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	type Input struct {
		target int
		result bool
	}
	testInput := []Input{
		{
			0,
			true,
		},
		{
			1,
			true,
		},
		{
			2,
			false,
		},
		{
			3,
			false,
		},
		{
			4,
			true,
		},
		{
			8,
			false,
		},
		{
			10,
			false,
		},
		{
			9,
			true,
		},
		{
			11,
			false,
		},
		{
			12,
			false,
		},
		{
			13,
			false,
		},
		{
			14,
			false,
		},
		{
			15,
			false,
		},
		{
			16,
			true,
		},
	}

	for _, input := range testInput {
		result := IsPerfectSquare(input.target)
		if result != input.result {
			t.Errorf("result != input.result, result = %v, input.result = %v, "+
				"target = %v", result, input.result, input.target)
		}
	}
}
