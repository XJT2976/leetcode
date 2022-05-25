package leetcode59

import "testing"

func TestGenerateMatrix(t *testing.T) {
	type Input struct {
		num    int
		expect [][]int
	}
	testInput := []Input{
		{
			1,
			[][]int{
				{1},
			},
		},
		{
			2,
			[][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			3,
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			4,
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			5,
			[][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	for _, input := range testInput {
		result := GenerateMatrix(input.num)
	loop:
		for i := 0; i < len(result); i++ {
			for j := 0; j < len(result[i]); j++ {
				if result[i][j] != input.expect[i][j] {
					t.Errorf("result != input.expect, result = %v, input.expect = %v", result, input.expect)
					break loop
				}
			}
		}
	}
}
