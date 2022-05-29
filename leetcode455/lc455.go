package leetcode455

import "sort"

func FindContentChildren(g []int, s []int) int {
	ret := 0
	sort.Ints(g)
	sort.Ints(s)
	for i, j := 0, 0; i < len(g) && j < len(s); {
		if s[j] >= g[i] {
			ret++
			i++
			j++
		} else {
			j++
		}
	}

	return ret
}
