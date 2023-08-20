package main

// https://space.bilibili.com/206214
func maximizeTheProfit(n int, offers [][]int) int {
	type pair struct{ start, gold int }
	groups := make([][]pair, n)
	for _, offer := range offers {
		start, end, gold := offer[0], offer[1], offer[2]
		groups[end] = append(groups[end], pair{start, gold})
	}

	f := make([]int, n+1)
	for end, g := range groups {
		f[end+1] = f[end]
		for _, p := range g {
			f[end+1] = max(f[end+1], f[p.start]+p.gold)
		}
	}
	return f[n]
}

func max(a, b int) int { if b > a { return b }; return a }
