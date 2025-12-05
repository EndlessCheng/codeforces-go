package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf78E(in io.Reader, out io.Writer) {
	var n, t int
	Fscan(in, &n, &t)
	a := make([]string, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	st := n * n * 2
	end := st + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, end+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}

	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	bfs0 := func(sx, sy int) [][]int {
		dis := make([][]int, n)
		for i := range dis {
			dis[i] = make([]int, n)
			for j := range dis[i] {
				dis[i][j] = 1e9
			}
		}
		dis[sx][sy] = 0
		q := []pair{{sx, sy}}
		for step := 1; len(q) > 0; step++ {
			tmp := q
			q = nil
			for _, p := range tmp {
				for _, d := range dir4 {
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < n && 0 <= y && y < n && a[x][y] <= '9' {
						if dis[x][y] == 1e9 {
							dis[x][y] = step
							q = append(q, pair{x, y})
						}
					}
				}
			}
		}
		return dis
	}

	var disZ [][]int
	ds := make([][][][]int, n)
	for i, r := range a {
		ds[i] = make([][][]int, n)
		for j, c := range r {
			if c == 'Z' {
				disZ = bfs0(i, j)
			} else if '0' < c && c <= '9' {
				addEdge(st, i*n+j, int(c-'0'))
				ds[i][j] = bfs0(i, j)
			}
		}
	}

	for i := range n {
		var s string
		Fscan(in, &s)
		for j, c := range s {
			if !('0' < c && c <= '9') {
				continue
			}
			capsule := n*n + i*n + j
			addEdge(capsule, end, int(c-'0'))
			for i0, r := range ds {
			o:
				for j0, d := range r {
					if d == nil || d[i][j] > t || d[i][j] > disZ[i][j] {
						continue
					}
					if i0 == i && j0 == j { // 这里忘判了 WA 了一发
						addEdge(i0*n+j0, capsule, 9)
						continue
					}
					for _, dir := range dir4 {
						x, y := i+dir.x, j+dir.y
						if 0 <= x && x < n && 0 <= y && y < n && a[x][y] <= '9' && d[x][y] < disZ[x][y] {
							addEdge(i0*n+j0, capsule, 9)
							continue o
						}
					}
				}
			}
		}
	}

	dis := make([]int, len(g))
	bfs := func() bool {
		clear(dis)
		dis[st] = 1
		q := []int{st}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && dis[w] == 0 {
					dis[w] = dis[v] + 1
					q = append(q, w)
				}
			}
		}
		return dis[end] > 0
	}
	iter := make([]int, len(g))
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == end {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && dis[w] > dis[v] {
				f := dfs(w, min(totalFlow-curFlow, e.cap))
				if f == 0 {
					continue
				}
				e.cap -= f
				g[w][e.rid].cap += f
				curFlow += f
				if curFlow == totalFlow {
					break
				}
			}
		}
		return
	}
	maxFlow := 0
	for bfs() {
		clear(iter)
		maxFlow += dfs(st, math.MaxInt)
	}
	Fprint(out, maxFlow)
}

//func main() { cf78E(bufio.NewReader(os.Stdin), os.Stdout) }
