package leetcode452

import "sort"

func FindMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}

		return points[i][0] < points[j][0]
	})

	arrowNum := 1
	r := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > r {
			arrowNum++
			r = points[i][1]
		} else {
			if points[i][1] < r {
				r = points[i][1]
			}
		}
	}

	return arrowNum
}
