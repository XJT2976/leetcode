package leetcode34

func SearchRange(nums []int, target int) []int {
	start := len(nums)
	end := -1

	l := 0
	h := start - 1

	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			start = mid
			end = mid
			for i := mid - 1; i >= l; i-- {
				if nums[i] == target {
					start = i
				} else {
					break
				}
			}
			for i := mid + 1; i <= h; i++ {
				if nums[i] == target {
					end = i
				} else {
					break
				}
			}
			break
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}

	if start > end {
		start = -1
		end = -1
	}
	return []int{start, end}
}
