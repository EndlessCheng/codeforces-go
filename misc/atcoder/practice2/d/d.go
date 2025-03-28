package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	dir4 := []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	var n, m int
	Fscan(in, &n, &m)
	a := make([][]byte, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	S := n * m
	T := S + 1
	type nb struct{ to, rid, cap int }
	g := make([][]nb, T+1)
	addEdge := func(from, to, cap int) {
		g[from] = append(g[from], nb{to, len(g[to]), cap})
		g[to] = append(g[to], nb{from, len(g[from]) - 1, 0})
	}
	for i, row := range a {
		for j := i % 2; j < m; j += 2 {
			if row[j] == '#' {
				continue
			}
			for _, dir := range dir4 {
				x, y := i+dir.x, j+dir.y
				if 0 <= x && x < n && 0 <= y && y < m && a[x][y] != '#' {
					addEdge(i*m+j, x*m+y, 1)
				}
			}
			addEdge(S, i*m+j, 1)
		}
		for j := i%2 ^ 1; j < m; j += 2 {
			if row[j] != '#' {
				addEdge(i*m+j, T, 1)
			}
		}
	}

	d := make([]int, len(g))
	bfs := func() bool {
		clear(d)
		d[S] = 1
		q := []int{S}
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, e := range g[v] {
				if w := e.to; e.cap > 0 && d[w] == 0 {
					d[w] = d[v] + 1
					q = append(q, w)
				}
			}
		}
		return d[T] > 0
	}
	iter := make([]int, len(g))
	var dfs func(int, int) int
	dfs = func(v, totalFlow int) (curFlow int) {
		if v == T {
			return totalFlow
		}
		for ; iter[v] < len(g[v]); iter[v]++ {
			e := &g[v][iter[v]]
			if w := e.to; e.cap > 0 && d[w] > d[v] {
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
		maxFlow += dfs(S, math.MaxInt)
	}
	Fprintln(out, maxFlow)

	img := [4]struct{ x, y byte }{{'<', '>'}, {'>', '<'}, {'^', 'v'}, {'v', '^'}}
	for i, row := range a {
		for j := i % 2; j < m; j += 2 {
			if row[j] == '#' {
				continue
			}
			cnt := 0
			for di, dir := range dir4 {
				x, y := i+dir.x, j+dir.y
				if 0 <= x && x < n && 0 <= y && y < m && a[x][y] != '#' {
					if g[i*m+j][cnt].cap == 0 {
						a[i][j] = img[di].x
						a[x][y] = img[di].y
					}
					cnt++
				}
			}
		}
	}
	for _, row := range a {
		Fprintf(out, "%s\n", row)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
