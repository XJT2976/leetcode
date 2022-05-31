package leetcode56

import "sort"

func Merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[i-1][1] {
			res = append(res, intervals[i-1])
		} else {
			if intervals[i-1][1] > intervals[i][1] {
				intervals[i][1] = intervals[i-1][1]
			}
			intervals[i][0] = intervals[i-1][0]
		}
	}

	res = append(res, intervals[len(intervals)-1])
	return res
}
