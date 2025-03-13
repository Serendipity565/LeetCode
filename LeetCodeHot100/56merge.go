package LeetCodeHot100

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	var ans [][]int
	l := intervals[0][0]
	r := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= r {
			r = max(r, intervals[i][1])
		} else {
			ans = append(ans, []int{l, r})
			l = intervals[i][0]
			r = intervals[i][1]
		}
	}
	if len(ans) == 0 {
		ans = append(ans, []int{l, r})
		return ans
	}
	if len(ans) != 0 && (l != ans[len(ans)-1][0] || r != ans[len(ans)-1][1]) {
		ans = append(ans, []int{l, r})
	}
	return ans
}
