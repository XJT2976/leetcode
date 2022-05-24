package leetcode69

import "testing"

func TestMySqrt(t *testing.T) {
	type Input struct {
		target int
		result int
	}
	testInput := []Input{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			1,
		},
		{
			3,
			1,
		},
		{
			4,
			2,
		},
		{
			8,
			2,
		},
		{
			10,
			3,
		},
		{
			9,
			3,
		},
		{
			11,
			3,
		},
		{
			12,
			3,
		},
		{
			13,
			3,
		},
		{
			14,
			3,
		},
		{
			15,
			3,
		},
		{
			16,
			4,
		},
	}

	for _, input := range testInput {
		result := MySqrt(input.target)
		if result != input.result {
			t.Errorf("result != input.result, result = %v, input.result = %v, "+
				"target = %v", result, input.result, input.target)
		}
	}
}
