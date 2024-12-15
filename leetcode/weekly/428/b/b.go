package main

// https://space.bilibili.com/206214
type pair struct {
	to   string
	rate float64
}

func calcAmount(pairs [][]string, rates []float64, initialCurrency string) map[string]float64 {
	g := map[string][]pair{}
	for i, p := range pairs {
		x, y, r := p[0], p[1], rates[i]
		g[x] = append(g[x], pair{y, r})
		g[y] = append(g[y], pair{x, 1 / r})
	}

	amount := map[string]float64{}
	var dfs func(string, float64)
	dfs = func(x string, curAmount float64) {
		amount[x] = curAmount
		for _, e := range g[x] {
			// 每个节点只需递归一次（重复递归算出来的结果是一样的，因为题目保证汇率没有矛盾）
			if amount[e.to] == 0 {
				dfs(e.to, curAmount*e.rate)
			}
		}
	}
	dfs(initialCurrency, 1)
	return amount
}

func maxAmount(initialCurrency string, pairs1 [][]string, rates1 []float64, pairs2 [][]string, rates2 []float64) (ans float64) {
	day1Amount := calcAmount(pairs1, rates1, initialCurrency)
	day2Amount := calcAmount(pairs2, rates2, initialCurrency)
	for x, a2 := range day2Amount {
		ans = max(ans, day1Amount[x]/a2)
	}
	return
}
