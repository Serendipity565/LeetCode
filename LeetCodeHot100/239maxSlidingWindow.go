package LeetCodeHot100

func maxSlidingWindow(nums []int, k int) []int {
	var ans []int
	var q []int
	n := len(nums)
	if n < k {
		cnt := -10000
		for i := 0; i < n; i++ {
			cnt = max(cnt, nums[i])
		}
		ans = append(ans, cnt)
	} else {
		for i := 0; i < k; i++ {
			for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
		ans = append(ans, nums[q[0]])
		for i := k; i < n; i++ {
			for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
			for q[0] <= i-k {
				q = q[1:]
			}
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
