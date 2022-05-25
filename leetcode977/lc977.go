package leetcode977

func SortedSquares(nums []int) []int {
	ret := make([]int, len(nums))
	l := 0
	h := len(nums) - 1
	pos := h

	for {
		if l > h {
			break
		}
		if abs(nums[l]) < abs(nums[h]) {
			ret[pos] = nums[h] * nums[h]
			h--
		} else {
			ret[pos] = nums[l] * nums[l]
			l++
		}

		pos--
	}

	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func SortedSquares2(nums []int) []int {
	pos := -1
	ret := make([]int, len(nums))
	for index, n := range nums {
		if n >= 0 {
			break
		}
		pos = index
	}

	l := pos
	r := pos + 1
	index := 0

	for {
		if l < 0 && r >= len(nums) {
			break
		}

		if l < 0 {
			for r < len(nums) {
				ret[index] = nums[r] * nums[r]
				r++
				index++
			}
			break
		}

		if r >= len(nums) {
			for l >= 0 {
				ret[index] = nums[l] * nums[l]
				l--
				index++
			}
			break
		}

		if -nums[l] > nums[r] {
			ret[index] = nums[r] * nums[r]
			r++
		} else {
			ret[index] = nums[l] * nums[l]
			l--
		}
		index++
	}

	return ret
}
