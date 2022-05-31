package leetcode135

func Candy(ratings []int) int {
	candySum := 0
	candy := make([]int, len(ratings))
	candy[0] = 1

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		} else {
			candy[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candy[i+1]+1 > candy[i] {
				candy[i] = candy[i+1] + 1
			}
		}
	}

	for i := range candy {
		candySum += candy[i]
	}
	return candySum
}
