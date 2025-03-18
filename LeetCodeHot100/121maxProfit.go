package LeetCodeHot100

func maxProfit(prices []int) int {
	pre := make([]int, len(prices))
	beh := make([]int, len(prices))
	pre[0] = prices[0]
	for i := 1; i < len(prices); i++ {
		pre[i] = min(pre[i-1], prices[i])
	}
	beh[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		beh[i] = max(beh[i+1], prices[i])
	}
	ans := 0
	for i := 0; i < len(prices); i++ {
		ans = max(ans, beh[i])
	}
	return ans
}
