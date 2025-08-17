package main

// https://space.bilibili.com/206214
func maxProfit1(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	sum := make([]int, n+1)
	sumSell := make([]int, n+1)
	for i, p := range prices {
		sum[i+1] = sum[i] + p*strategy[i]
		sumSell[i+1] = sumSell[i] + p
	}

	ans := sum[n] // 不修改
	for i := k; i <= n; i++ {
		res := sum[i-k] + sum[n] - sum[i] + sumSell[i] - sumSell[i-k/2]
		ans = max(ans, res)
	}
	return int64(ans)
}

func maxProfit(prices, strategy []int, k int) int64 {
	total, sum := 0, 0
	// 计算第一个窗口
	for i := range k / 2 {
		p, s := prices[i], strategy[i]
		total += p * s
		sum -= p * s
	}
	for i := k / 2; i < k; i++ {
		p, s := prices[i], strategy[i]
		total += p * s
		sum += p * (1 - s)
	}
	maxSum := max(sum, 0)

	for i := k; i < len(prices); i++ {
		p, s := prices[i], strategy[i]
		total += p * s
		sum += p*(1-s) - prices[i-k/2] + prices[i-k]*strategy[i-k]
		maxSum = max(maxSum, sum)
	}
	return int64(total + maxSum)
}
