package leetcode22

func GenerateParenthesis(n int) []string {
	m := map[byte]int{}
	ret := []string{}
	m['('], m[')'] = n, n
	o := make([]byte, 2*n)
	dfs(o, m, 0, &ret)
	return ret
}

func dfs(arr []byte, m map[byte]int, pos int, ret *[]string) {
	if m['('] == 0 && m[')'] == 0 {
		*ret = append(*ret, string(arr))
		return
	}
	if m['('] > 0 {
		arr[pos] = '('
		m['('] -= 1
		pos += 1
		dfs(arr, m, pos, ret)
		m['('] += 1
		pos -= 1
	}

	if m[')'] > 0 && m[')'] > m['('] {
		arr[pos] = ')'
		m[')'] -= 1
		pos += 1
		dfs(arr, m, pos, ret)
		m[')'] += 1
		pos -= 1
	}
}
