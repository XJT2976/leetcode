package leetcode90

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)
	backtracking(0, nums, &path, &res)
	return res
}

func backtracking(start int, nums []int, path *[]int, result *[][]int) {
	tmp := make([]int, len(*path))
	copy(tmp, *path)
	*result = append(*result, tmp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*path = append(*path, nums[i])
		backtracking(i+1, nums, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
