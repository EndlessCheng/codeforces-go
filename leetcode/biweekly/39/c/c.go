package main

// github.com/EndlessCheng/codeforces-go
func minimumJumps(forbidden []int, a, b, x int) (ans int) {
	fb := map[int]bool{}
	for _, v := range forbidden {
		fb[v] = true
	}
	const mx int = 1e4
	vis := [mx][2]bool{}
	q := [][2]int{{}}
	for ; len(q) > 0; ans++ {
		qq := q
		q = nil
		for _, p := range qq {
			v := p[0]
			if v == x {
				return
			}
			if p[1] > 0 && v-b > 0 && !vis[v-b][0] && !fb[v-b] {
				vis[v-b][0] = true
				q = append(q, [2]int{v - b, 0})
			}
			if v+a < mx && !vis[v+a][1] && !fb[v+a] {
				vis[v+a][1] = true
				q = append(q, [2]int{v + a, 1})
			}
		}
	}
	return -1
}
