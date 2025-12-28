package main

// https://space.bilibili.com/206214
func minimumCost(cost1, cost2, costBoth, need1, need2 int) int64 {
	res1 := cost1*need1 + cost2*need2 // 各买各的
	if need1 > need2 {
		need1, need2 = need2, need1
		cost2 = cost1
	}
	res2 := costBoth * need2 // 我包了
	res3 := costBoth*need1 + cost2*(need2-need1) // 混合策略
	return int64(min(res1, res2, res3))
}
