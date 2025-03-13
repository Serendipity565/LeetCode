package LeetCodeHot100

func spiralOrder(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	ans := make([]int, (right+1)*(bottom+1))
	now := 0
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			ans[now] = matrix[top][i]
			now++
			if now == len(ans) {
				return ans
			}
		}
		for i := top + 1; i < bottom; i++ {
			ans[now] = matrix[i][right]
			now++
			if now == len(ans) {
				return ans
			}
		}
		for i := right; i >= left; i-- {
			ans[now] = matrix[bottom][i]
			now++
			if now == len(ans) {
				return ans
			}
		}
		for i := bottom - 1; i > top; i-- {
			ans[now] = matrix[i][left]
			now++
			if now == len(ans) {
				return ans
			}
		}
		top++
		bottom--
		left++
		right--
	}
	return ans
}
