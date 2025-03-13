package LeetCodeHot100

func maxSubArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	left := 0
	var ans = nums[0]
	for rihgt := 1; rihgt < len(nums); rihgt++ {
		left = min(left, nums[rihgt-1])
		ans = max(ans, nums[rihgt]-left)
	}
	return ans
}
