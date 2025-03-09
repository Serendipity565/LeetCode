package LeetCodeHot100

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	has := map[int]bool{}
	for _, num := range nums {
		has[num] = true
	}
	ans := 1
	for x := range has {
		if !has[x-1] && has[x] {
			now := x
			nowlen := 1
			has[now] = false
			for y := x + 1; has[y]; y++ {
				nowlen++
				has[y] = false
			}
			if nowlen > ans {
				ans = nowlen
			}
		}
	}
	return ans
}
