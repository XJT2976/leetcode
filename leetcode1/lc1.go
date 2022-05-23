package leetcode1

func TwoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for index, num := range nums {
		if i, ok := indexMap[target-num]; ok {
			return []int{index, i}
		}

		indexMap[num] = index
	}
	return nil
}
