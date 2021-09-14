package main

func countSubTrees(n int, edges [][]int, s string) (ans []int) {
	g := make([][]int, n)
	for _, e := range edges {
		v, w := e[0], e[1]
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	ans = make([]int, n)
	var f func(v, fa int) [26]int
	f = func(v, fa int) (cnt [26]int) {
		cnt[s[v]-'a'] = 1
		for _, w := range g[v] {
			if w != fa {
				for i, c := range f(w, v) {
					cnt[i] += c
				}
			}
		}
		ans[v] = cnt[s[v]-'a']
		return
	}
	f(0, -1)
	return
}
