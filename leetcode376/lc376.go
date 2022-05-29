package leetcode376

func WiggleMaxLength(nums []int) int {
	const (
		positive = iota
		equal
		negative
	)
	if len(nums) == 0 {
		return 0
	}

	state := equal
	ret := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if state == negative || state == equal {
				state = positive
				ret += 1
			}
		} else if nums[i] < nums[i-1] {
			if state == positive || state == equal {
				state = negative
				ret += 1
			}
		}
	}

	return ret
}
