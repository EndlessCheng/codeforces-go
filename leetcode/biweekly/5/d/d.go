package main

// github.com/EndlessCheng/codeforces-go
func minimumSemesters(n int, relations [][]int) (ans int) {
	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range relations {
		v, w := e[0]-1, e[1]-1
		g[v] = append(g[v], w)
		deg[w]++
	}

	levels := make([]int, n)
	cnt := 0
	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
			levels[i] = 1
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		cnt++
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				levels[w] = levels[v] + 1
				ans = levels[w]
				q = append(q, w)
			}
		}
	}
	if cnt < n {
		return -1
	}
	return
}
