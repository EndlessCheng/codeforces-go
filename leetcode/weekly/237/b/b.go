package main

import "slices"

// github.com/EndlessCheng/codeforces-go
func maxIceCream(costs []int, coins int) int {
	slices.Sort(costs)

	// 按照价格从低到高买
	for i, cost := range costs {
		if coins < cost { // 钱不够
			return i // 买 [0, i-1] 一共 i 根雪糕
		}
		coins -= cost
	}

	// 可以买所有雪糕
	return len(costs)
}

func maxIceCream2(costs []int, coins int) (ans int) {
	mx := slices.Max(costs)
	cnt := make([]int, mx+1)
	for _, cost := range costs {
		cnt[cost]++
	}

	// 按照价格从低到高买
	for cost := 1; cost <= mx && cost <= coins; cost++ {
		num := min(cnt[cost], coins/cost)
		coins -= cost * num // 买 num 根雪糕
		ans += num
	}
	return
}
