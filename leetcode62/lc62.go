package leetcode62

func UniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		if len(dp[i]) == 0 {
			dp[i] = make([]int, n)
		}

		for j := range dp[i] {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}