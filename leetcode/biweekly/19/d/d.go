package main

// github.com/EndlessCheng/codeforces-go
func minJumps(a []int) int {
	pos := map[int][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}
	n := len(a)
	vis := make([]bool, n)
	vis[0] = true
	type pair struct{ v, d int }
	q := []pair{{0, 0}}
	for {
		p := q[0]
		q = q[1:]
		v, d := p.v, p.d
		if v == n-1 {
			return d
		}
		if v > 0 && !vis[v-1] {
			vis[v-1] = true
			q = append(q, pair{v - 1, d + 1})
		}
		if !vis[v+1] {
			vis[v+1] = true
			q = append(q, pair{v + 1, d + 1})
		}
		if pos[a[v]] == nil {
			continue
		}
		for _, w := range pos[a[v]] {
			if !vis[w] {
				vis[w] = true
				q = append(q, pair{w, d + 1})
			}
		}
		delete(pos, a[v])
	}
}
