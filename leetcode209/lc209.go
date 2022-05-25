package leetcode209

func MinSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	l := 0
	r := 0
	if len(nums) <= 0 {
		return 0
	}
	sum := nums[0]

	for {
		if sum >= target {
			if r-l+1 < minLen {
				minLen = r - l + 1
			}
			sum -= nums[l]
			l++
			if l > r {
				break
			}
		} else {
			r++
			if r >= len(nums) {
				break
			}
			sum += nums[r]
		}
	}

	if minLen >= len(nums)+1 {
		return 0
	} else {
		return minLen
	}
}
