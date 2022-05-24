package leetcode704

func Search(nums []int, target int) int {
	l := 0
	h := len(nums) - 1

	for l <= h {
		mid := (h-l)/2 + l
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			h = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func Search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if target > nums[len(nums)/2-1] {
		if index := Search(nums[len(nums)/2:], target); index != -1 {
			return index + len(nums)/2
		} else {
			return -1
		}
	} else {
		return Search(nums[0:len(nums)/2], target)
	}
}
