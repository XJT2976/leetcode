package leetcode367

func IsPerfectSquare(num int) bool {
	l := 0
	h := num

	if num == 0 || num == 1 {
		return true
	}
	for l <= h {
		mid := (h + l) / 2
		if mid == l {
			return false
		}
		sq := mid * mid
		if sq == num {
			return true
		} else if sq > num {
			h = mid
		} else {
			l = mid
		}
	}

	return false
}
