package main

// github.com/EndlessCheng/codeforces-go
func regionsBySlashes(a []string) (ans int) {
	n := len(a)
	m := n * n
	g := make([][]int, 2*m)
	add := func(v, w int) {
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for i, r := range a {
		for j, b := range r {
			v := i*n + j
			w := v + m
			if i > 0 {
				add(v, w-n)
			}
			if i < n-1 {
				add(w, v+n)
			}
			if b == '/' {
				if j > 0 {
					if r[j-1] == '/' {
						add(v, w-1)
					} else {
						add(v, v-1)
					}
				}
				if j < n-1 {
					if r[j+1] == '/' {
						add(w, v+1)
					} else {
						add(w, w+1)
					}
				}
			} else {
				if b == ' ' {
					add(v, w)
				}
				if j > 0 {
					if r[j-1] == '/' {
						add(w, w-1)
					} else {
						add(w, v-1)
					}
				}
				if j < n-1 {
					if r[j+1] == '/' {
						add(v, v+1)
					} else {
						add(v, w+1)
					}
				}
			}
		}
	}

	vis := make([]bool, len(g))
	var f func(int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, b := range vis {
		if !b {
			ans++
			f(i)
		}
	}
	return
}
