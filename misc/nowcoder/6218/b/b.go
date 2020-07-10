package main

// github.com/EndlessCheng/codeforces-go
func solve(n int, m int) int {
	const mx int = 2e3
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
		ws := []int{}
		if v < mx {
			ws = append(ws, v+1)
		}
		if v > 1 {
			ws = append(ws, v-1)
		}
		if v*v <= mx {
			ws = append(ws, v*v)
		}
		for _, w := range ws {
			if !vis[w] {
				vis[w] = true
				q = append(q, pair{w, d + 1})
			}
		}
	}
}
