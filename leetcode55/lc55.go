package leetcode55

func CanJump(nums []int) bool {
	for i := 0; ; {
		if i >= len(nums)-1 || nums[i]+i >= len(nums)-1 {
			break
		}

		if nums[i] == 0 {
			return false
		}

		maxStep := i + 1 + nums[i+1]
		nextI := i + 1
		for j := i + 2; j <= i+nums[i]; j++ {
			if j+nums[j] >= maxStep {
				maxStep = j + nums[j]
				nextI = j
			}
		}

		i = nextI
	}

	return true
}
