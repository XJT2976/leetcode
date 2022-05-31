package leetcode435

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	eraseNum := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[i-1][1] {
			continue
		} else {
			if intervals[i][1] > intervals[i-1][1] {
				intervals[i][1] = intervals[i-1][1]
			}
			eraseNum++
		}
	}

	return eraseNum
}
