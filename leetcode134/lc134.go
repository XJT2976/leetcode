package leetcode134

func CanCompleteCircuit(gas []int, cost []int) int {
	rest := make([]int, len(gas))
	sum := 0
	for i := range rest {
		rest[i] = gas[i] - cost[i]
		sum += rest[i]
	}

	if sum < 0 {
		return -1
	}

	sum = 0
	startIndex := -1
	for i := range rest {
		sum += rest[i]
		if sum < 0 {
			startIndex = -1
			sum = 0
		} else if startIndex == -1 {
			startIndex = i
		}
	}

	return startIndex
}
