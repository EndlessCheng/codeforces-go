package main

// https://space.bilibili.com/206214
func isPossible(n int, edges [][]int) bool {
	type pair struct{ x, y int }
	has := map[pair]bool{}
	deg := make([]int, n+1)
	for _, e := range edges {
		x, y := e[0], e[1]
		has[pair{x, y}] = true
		has[pair{y, x}] = true
		deg[x]++
		deg[y]++
	}
	odd := []int{}
	for i, d := range deg {
		if d%2 > 0 {
			odd = append(odd, i)
		}
	}
	m := len(odd)
	if m == 0 {
		return true
	}
	if m == 2 {
		x, y := odd[0], odd[1]
		if !has[pair{x, y}] {
			return true
		}
		for i := 1; i <= n; i++ {
			if i != x && i != y && !has[pair{i, x}] && !has[pair{i, y}] {
				return true
			}
		}
		return false
	}
	if m == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return !has[pair{a, b}] && !has[pair{c, d}] || !has[pair{a, c}] && !has[pair{b, d}] || !has[pair{a, d}] && !has[pair{b, c}]
	}
	return false
}
