package main

func isPathCrossing(s string) (ans bool) {
	type pair struct{ x, y int }
	dir4C := [...]pair{
		'W': {-1, 0},
		'E': {1, 0},
		'S': {0, -1},
		'N': {0, 1},
	}
	vis := map[pair]bool{{}: true}
	p := pair{}
	for _, c := range s {
		d := dir4C[c]
		p.x += d.x
		p.y += d.y
		if vis[p] {
			return true
		}
		vis[p] = true
	}
	return
}
