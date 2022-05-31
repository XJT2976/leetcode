package leetcode1005

import "sort"

func LargestSumAfterKNegations(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < k; i++ {
		nums[0] = -nums[0]
		sort.Ints(nums)
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	return sum
}
