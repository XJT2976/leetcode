package leetcode2150

import "sort"

func FindLonely(nums []int) []int {
	res := make([]int, 0, len(nums))
	sort.Ints(nums)
	for i := range nums {
		if (i > 0 && ((nums[i-1] == nums[i]) || (nums[i-1] == (nums[i] - 1)))) ||
			(i < len(nums)-1 && ((nums[i+1] == nums[i]) || (nums[i+1] == (nums[i] + 1)))) {
			continue
		}
		res = append(res, nums[i])
	}

	return res
}
