package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func findMaximumElegance(items [][]int, k int) int64 {
	// 把利润从大到小排序
	slices.SortFunc(items, func(a, b []int) int { return b[0] - a[0] })
	ans, totalProfit := 0, 0
	vis := map[int]bool{}
	duplicate := []int{} // 栈
	for i, p := range items {
		profit, category := p[0], p[1]
		if i < k {
			totalProfit += profit
			if !vis[category] {
				vis[category] = true
			} else { // 重复类别
				duplicate = append(duplicate, profit)
			}
		} else if len(duplicate) > 0 && !vis[category] {
			vis[category] = true
			totalProfit += profit - duplicate[len(duplicate)-1] // 选一个重复类别中的最小利润替换
			duplicate = duplicate[:len(duplicate)-1]
		} // else：比前面的利润小，而且类别还重复了，选它只会让 totalProfit 变小，len(vis) 不变，优雅度不会变大
		ans = max(ans, totalProfit+len(vis)*len(vis))
	}
	return int64(ans)
}
