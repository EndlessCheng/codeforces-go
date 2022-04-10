package main

// github.com/EndlessCheng/codeforces-go
func lightSticks(n int, m int, indices []int) (ans []int) {
	tot := n*m*2 + n + m
	del := make([]bool, tot)
	for _, i := range indices {
		del[i] = true
	}

	mi := int(1e9)
	mm := 2*m + 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			vis := append([]bool{}, del...)
			left := tot - len(indices)
			k := 0
			type pair struct{ x, y int }
			q := []pair{{i, j}}
			for ; len(q) > 0; k++ {
				tmp := q
				q = nil
				for _, p := range tmp {
					x, y := p.x, p.y
					if id := (x-1)*mm + m + y; x > 0 && !vis[id] {
						vis[id] = true
						left--
						q = append(q, pair{x - 1, y})
					}
					if id := x*mm + m + y; x < n && !vis[id] {
						vis[id] = true
						left--
						q = append(q, pair{x + 1, y})
					}
					if id := x*mm + y - 1; y > 0 && !vis[id] {
						vis[id] = true
						left--
						q = append(q, pair{x, y - 1})
					}
					if id := x*mm + y; y < m && !vis[id] {
						vis[id] = true
						left--
						q = append(q, pair{x, y + 1})
					}
				}
			}
			if left > 0 {
				continue
			}
			if k < mi {
				mi = k
				ans = []int{i*(m+1) + j}
			} else if k == mi {
				ans = append(ans, i*(m+1)+j)
			}
		}
	}
	return
}
