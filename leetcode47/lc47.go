package leetcode47

func PermuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	backtracking(nums, used, &path, &res)
	return res
}

func backtracking(nums []int, used map[int]bool, path *[]int, result *[][]int) {
	if len(*path) == len(nums) {
		tmp := make([]int, len(*path))
		for i := range *path {
			tmp[i] = (*path)[i]
		}
		*result = append(*result, tmp)
	}

	usedThisLevel := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if used[i] || usedThisLevel[nums[i]] {
			continue
		}
		used[i] = true
		usedThisLevel[nums[i]] = true
		*path = append(*path, nums[i])
		backtracking(nums, used, path, result)
		*path = (*path)[:len(*path)-1]
		used[i] = false
	}
}
