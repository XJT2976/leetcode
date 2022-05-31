package leetcode45

func Jump(nums []int) int {
	jumpCount := 0
	if len(nums) <= 1 {
		return 0
	}

	for i := 0; ; {
		if i+nums[i] >= len(nums)-1 {
			jumpCount += 1
			break
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
		jumpCount += 1
	}

	return jumpCount
}
