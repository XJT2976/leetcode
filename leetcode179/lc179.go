package leetcode179

import (
	"sort"
	"strconv"
)

func LargestNumber(nums []int) string {
	// 360 3604
	// 360436

	sort.Slice(nums, func(i, j int) bool {
		si := strconv.Itoa(nums[i])
		sj := strconv.Itoa(nums[j])
		if si+sj > sj+si {
			return true
		} else {
			return false
		}
	})

	s := ""
	for i := range nums {
		s += strconv.Itoa(nums[i])
	}
	return s
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
