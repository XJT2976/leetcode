package leetcode860

func LemonadeChange(bills []int) bool {
	ownBills := make([]int, 11)

	for i := range bills {
		switch bills[i] {
		case 5:
			ownBills[5] += 1
		case 10:
			if ownBills[5] == 0 {
				return false
			}
			ownBills[5] -= 1
			ownBills[10] += 1
		case 20:
			if ownBills[10] >= 1 && ownBills[5] >= 1 {
				ownBills[5] -= 1
				ownBills[10] -= 1
			} else if ownBills[5] >= 3 {
				ownBills[5] -= 3
			} else {
				return false
			}
		}
	}

	return true
}
