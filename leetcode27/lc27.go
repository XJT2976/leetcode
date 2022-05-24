package leetcode27

func RemoveElement(nums []int, val int) int {
	l := 0
	h := 0
	k := 0

	for h < len(nums) {
		if nums[h] != val {
			nums[l] = nums[h]
			l++
			k++
		}
		h++
	}

	return k
}
