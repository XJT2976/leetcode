package leetcode491

func FindSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtracking(0, nums, &path, &res)
	return res
}

func backtracking(start int, nums []int, path *[]int, result *[][]int) {
	if len(*path) >= 2 {
		tmp := make([]int, len(*path))
		for i := range *path {
			tmp[i] = (*path)[i]
		}
		*result = append(*result, tmp)
	}

	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		if len(*path) > 0 && nums[i] < (*path)[len(*path)-1] {
			continue
		}
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		*path = append(*path, nums[i])
		backtracking(i+1, nums, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
