package LeetCodeHot100

func subarraySum(nums []int, k int) int {
	ans := 0

	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			ans++
			if v, ok := m[0]; ok {
				ans += v
			}
		} else {
			if v, ok := m[nums[i]-k]; ok {
				ans += v
			}
		}
		m[nums[i]]++
	}

	return ans
}
