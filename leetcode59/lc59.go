package leetcode59

func GenerateMatrix(n int) [][]int {
	const (
		left  = 0
		right = 1
		up    = 2
		down  = 3
	)
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	if n == 0 {
		return ret
	}

	leftBound := 0
	rightBound := n
	upBound := 1
	downBound := n
	direction := right
	ret[0][0] = 1
	x := 1
	y := 0

	for i := 2; i <= n*n; i++ {
		ret[y][x] = i
		switch direction {
		case left:
			if x <= leftBound {
				leftBound += 1
				direction = up
				y--
			} else {
				x--
			}
		case right:
			if x >= rightBound-1 {
				rightBound -= 1
				direction = down
				y++
			} else {
				x++
			}
		case up:
			if y <= upBound {
				upBound += 1
				direction = right
				x++
			} else {
				y--
			}
		case down:
			if y >= downBound-1 {
				downBound -= 1
				direction = left
				x--
			} else {
				y++
			}
		}
	}

	return ret
}
