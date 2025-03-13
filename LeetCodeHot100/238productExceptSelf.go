package LeetCodeHot100

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i] = nums[i] * nums[i+1]
	}
	ans[len(nums)-1] = ans[len(nums)-2]
	temp := ans[0]
	for i := 1; i < len(nums)-1; i++ {
		temp2 := ans[i]
		ans[i] = temp * nums[i+1]
		temp = temp2
	}
	ans[0] = nums[1]
	return ans
}
