package LeetCodeHot100

func rotate189(nums []int, k int) {
	n := len(nums)
	k %= n
	temp := append(nums[n-k:], nums[:n-k]...)
	copy(nums, temp)
}
