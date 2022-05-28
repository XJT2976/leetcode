package leetcode131

func Partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	backtracking(0, 0, s, &path, &res)
	return res
}

func backtracking(start, end int, s string, path *[]string, result *[][]string) {
	if len(s) == start {
		if len(*path) != 0 {
			tmp := make([]string, len(*path))
			for i := range *path {
				tmp[i] = (*path)[i]
			}
			*result = append(*result, tmp)
		}

		return
	}

	// "cbbbcc"
	for end < len(s) {
		if isPalindrome(start, end, s) {
			*path = append(*path, s[start:end+1])
			backtracking(end+1, end+1, s, path, result)
			*path = (*path)[:len(*path)-1]
		}
		end++
	}
}

func isPalindrome(start, end int, s string) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
