package main

// github.com/EndlessCheng/codeforces-go
func isPrintable(a [][]int) (ans bool) {
	var mix, miy, mxx, mxy [61]int
	for i := 1; i < 61; i++ {
		mix[i] = 99
		miy[i] = 99
	}
	for i, row := range a {
		for j, v := range row {
			mix[v] = min(mix[v], i)
			miy[v] = min(miy[v], j)
			mxx[v] = max(mxx[v], i)
			mxy[v] = max(mxy[v], j)
		}
	}
	g := [61][]int{}
	deg := [61]int{-1}
	for v := 1; v < 60; v++ {
		if mix[v] == 99 {
			continue
		}
		for w := v + 1; w < 61; w++ {
			if mix[w] == 99 {
				continue
			}
			up, down, l, r := max(mix[v], mix[w]), min(mxx[v], mxx[w]), max(miy[v], miy[w]), min(mxy[v], mxy[w])
			if up > down || l > r {
				continue
			}
			hasV, hasW := false, false
			for _, row := range a[up : down+1] {
				for _, x := range row[l : r+1] {
					if x == v {
						if hasW {
							return
						}
						hasV = true
					} else if x == w {
						if hasV {
							return
						}
						hasW = true
					}
				}
			}
			if hasV {
				g[w] = append(g[w], v)
				deg[v]++
			} else if hasW {
				g[v] = append(g[v], w)
				deg[w]++
			}
		}
	}
	q := []int{}
	for i, d := range deg[:] {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, w := range g[v] {
			if deg[w]--; deg[w] == 0 {
				q = append(q, w)
			}
		}
	}
	for _, d := range deg[1:] {
		if d > 0 {
			return
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
