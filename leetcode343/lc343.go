package leetcode343

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
