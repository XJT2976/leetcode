package leetcode202

func IsHappy(n int) bool {
	m := make(map[int]bool)

	for {
		if n == 1 {
			break
		}
		if _, ok := m[n]; ok {
			return false
		}

		m[n] = true
		sum := 0
		for n/10 > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		sum += (n % 10) * (n % 10)
		n = sum
	}

	return true
}
