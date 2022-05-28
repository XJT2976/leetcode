package leetcode93

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)
	backtracking(0, 0, s, &path, &res)
	return res
}

func backtracking(start, end int, s string, path *[]string, result *[]string) {
	if len(*path) > 4 {
		return
	}
	if len(s) == start {
		if len(*path) == 4 {
			tmp := strings.Join(*path, ".")
			*result = append(*result, tmp)
		}

		return
	}

	for end < len(s) {
		if end-start >= 3 {
			break
		}
		if isValidIp(start, end, s) {
			*path = append(*path, s[start:end+1])
			backtracking(end+1, end+1, s, path, result)
			*path = (*path)[:len(*path)-1]
		}
		end++
	}
}

func isValidIp(start, end int, s string) bool {
	if end-start >= 3 {
		return false
	}

	if s[start] == '0' && start != end {
		return false
	}

	i, _ := strconv.Atoi(s[start : end+1])
	if i > 255 {
		return false
	}

	return true
}
