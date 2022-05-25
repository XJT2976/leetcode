package leetcode242

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[int32]int)
	for _, c := range s {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}

	for _, c := range t {
		count, ok := m[c]
		if !ok || count <= 0 {
			return false
		}

		m[c] -= 1
	}

	return true
}
