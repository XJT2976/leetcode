package leetcode39

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)
	backtracking(candidates, target, 0, 0, &path, &res)
	return res
}

func backtracking(candidates []int, target, sum, start int, path *[]int, result *[][]int) {
	if sum == target {
		if len(*path) != 0 {
			tmp := make([]int, len(*path))
			for i := range *path {
				tmp[i] = (*path)[i]
			}
			*result = append(*result, tmp)
		}

		return
	}

	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		if sum > target {
			break
		}
		*path = append(*path, candidates[i])
		backtracking(candidates, target, sum, i, path, result)
		*path = (*path)[:len(*path)-1]
		sum -= candidates[i]
	}
}
