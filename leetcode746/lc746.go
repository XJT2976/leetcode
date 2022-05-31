package leetcode746

import "fmt"

func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}

	fmt.Println(dp)

	return dp[len(cost)]
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
