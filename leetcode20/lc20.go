package leetcode20

func IsValid(s string) bool {
	stack := make([]uint8, len(s))
	j := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			stack[j] = s[i]
			j++
		case ')':
			if j <= 0 || stack[j-1] != '(' {
				return false
			}
			j--
		case ']':
			if j <= 0 || stack[j-1] != '[' {
				return false
			}
			j--
		case '}':
			if j <= 0 || stack[j-1] != '{' {
				return false
			}
			j--
		}
	}

	if j > 0 {
		return false
	}
	return true
}
