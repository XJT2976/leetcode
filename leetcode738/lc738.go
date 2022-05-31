package leetcode738

import "strconv"

func MonotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	index := len(s)

	for i := len(s) - 2; i >= 0; i-- {
		if s[i] > s[i+1] {
			index = i + 1
			s[i] -= 1
		}
	}

	for i := index; i < len(s); i++ {
		s[i] = '9'
	}

	res, _ := strconv.Atoi(string(s))
	return res
}
