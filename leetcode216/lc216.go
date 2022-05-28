package leetcode216

func CombinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtracking(1, 9, k, n, 0, &path, &res)
	return res
}

func backtracking(start, end, k, n, sum int, path *[]int, result *[][]int) {
	if len(*path) == k {
		if sum == n {
			tmp := make([]int, len(*path))
			for i := range *path {
				tmp[i] = (*path)[i]
			}
			*result = append(*result, tmp)
		}

		return
	}

	for i := start; i <= end; i++ {
		sum += i
		if sum > n {
			break
		}
		*path = append(*path, i)
		backtracking(i+1, end, k, n, sum, path, result)
		*path = (*path)[:len(*path)-1]
		sum -= i
	}
}
