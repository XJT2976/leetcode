package leetcode69

func MySqrt(x int) int {
	target := x
	l := 0
	h := x

	if x == 0 || x == 1 {
		return x
	}

	for h >= l {
		mid := (h + l) / 2
		sq := mid * mid
		if sq == target {
			return mid
		} else if sq < target {
			if h-l <= 1 {
				return l
			}
			l = mid
		} else {
			if h-l <= 1 {
				return l
			}
			h = mid
		}
	}

	return 0
}
