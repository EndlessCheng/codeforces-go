package main

const mx = 2000
var g = [...][]int{1: {2}, mx: {mx - 1}}

func init() {
	for i := 2; i < mx; i++ {
		if i*i <= mx {
			g[i] = append(g[i], i*i)
		}
		g[i] = append(g[i], i-1, i+1)
	}
}

// github.com/EndlessCheng/codeforces-go
func solve(n int, m int) int {
	vis := [mx + 1]bool{}
	vis[n] = true
	type pair struct{ v, d int }
	q := []pair{{n, 0}}
	for {
		p := q[0]
		q = q[1:]
		v, d := p.v, p.d
		if v == m {
			return d
		}
		for _, w := range g[v] {
			if !vis[w] {
				vis[w] = true
				q = append(q, pair{w, d + 1})
			}
		}
	}
}
