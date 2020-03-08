package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) (ans int) {
	g := make([][]int, n)
	for w, v := range manager {
		if v != -1 {
			g[v] = append(g[v], w)
		}
	}
	var f func(v, t int)
	f = func(v, t int) {
		if len(g[v]) == 0 && t > ans {
			ans = t
		}
		for _, w := range g[v] {
			f(w, t+informTime[v])
		}
	}
	f(headID, 0)
	return
}
