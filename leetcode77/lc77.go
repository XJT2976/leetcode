package leetcode77

func Combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtracking(1, n, k, &path, &res)
	return res
}

func backtracking(pos, end, k int, path *[]int, result *[][]int) {
	for end-pos+1 >= k {
		*path = append(*path, pos)
		if k == 1 {
			if pos <= end {
				tmp := make([]int, len(*path))
				for i := range *path {
					tmp[i] = (*path)[i]
				}
				*result = append(*result, tmp)
			}
		} else {
			backtracking(pos+1, end, k-1, path, result)
		}
		*path = (*path)[:len(*path)-1]
		pos += 1
	}
}
