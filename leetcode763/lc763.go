package leetcode763

func PartitionLabels(s string) []int {
	bound := make([]int, 'z'+1)
	res := make([]int, 0)
	for i := range s {
		bound[s[i]] = i
	}

	left := 0
	right := 0
	for i := range s {
		if bound[s[i]] > right {
			right = bound[s[i]]
		}
		if i == right {
			res = append(res, right-left+1)
			left = right + 1
		}
	}

	return res
}
