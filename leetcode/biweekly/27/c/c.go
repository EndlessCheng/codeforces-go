package main

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) (ans []bool) {
	g := make([][]bool, n)
	for i := range g {
		g[i] = make([]bool, n)
	}
	for _, e := range prerequisites {
		g[e[0]][e[1]] = true
	}
	for k := range g {
		for i := range g {
			for j := range g {
				g[i][j] = g[i][j] || g[i][k] && g[k][j]
			}
		}
	}
	for _, q := range queries {
		ans = append(ans, g[q[0]][q[1]])
	}
	return
}
