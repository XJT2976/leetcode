package leetcode17

var numLetterTab = [][]uint8{
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

func LetterCombinations(digits string) []string {
	res := make([]string, 0)
	path := make([]uint8, 0)
	backtracking(digits, &path, &res)
	return res
}

func backtracking(digits string, path *[]uint8, result *[]string) {
	if len(digits) == 0 {
		if len(*path) != 0 {
			*result = append(*result, string(*path))
		}
		return
	}

	for i := 0; i < len(numLetterTab[digits[0]-'2']); i++ {
		*path = append(*path, numLetterTab[digits[0]-'2'][i])
		backtracking(digits[1:], path, result)
		*path = (*path)[:len(*path)-1]
	}
}
