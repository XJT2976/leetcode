package leetcode454

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum := 0
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			m[nums1[i]+nums2[j]] += 1
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum += m[-(nums3[i] + nums4[j])]
		}
	}

	return sum
}
