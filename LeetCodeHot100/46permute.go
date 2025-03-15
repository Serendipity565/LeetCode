package LeetCodeHot100

func permute(nums []int) [][]int {
	check := make([]bool, len(nums))
	ans := make([]int, len(nums))
	var res [][]int
	var dfs func(i int)
	dfs = func(l int) {
		if l == len(nums) {
			res = append(res, append([]int(nil), ans...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if !check[j] {
				check[j] = true
				ans[l] = nums[j]
				dfs(l + 1)
				check[j] = false
			}
		}
	}
	dfs(0)
	return res
}
