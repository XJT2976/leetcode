package leetcode1512

func NumIdenticalPairs(nums []int) int {
    sum := 0
	m := map[int]int{}
	for _, i := range nums {
		m[i] += 1
	}

	for _, v := range m {
		// 1 + 2 + ... + n - 1
		// (1 + n - 1) * (n - 1) / 2 = (n * n - n) / 2
		if v >= 2 {
			sum += (v * v - v) / 2
		}
	}

	return sum
}