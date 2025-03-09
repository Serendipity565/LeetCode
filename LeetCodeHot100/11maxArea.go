package LeetCodeHot100

func maxArea(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	left[0] = left[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}
	left[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	mymax := min(left[0], right[n-1]) * (n - 1)
	i := 0
	j := n - 1
	for i < j {
		if left[i] < right[j] {
			i++
			mymax = max(mymax, min(left[i], right[j])*(j-i))
		} else {
			j--
			mymax = max(mymax, min(left[i], right[j])*(j-i))
		}
	}
	return mymax
}
