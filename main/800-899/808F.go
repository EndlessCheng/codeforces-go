package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF808F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e9
	type card struct{ p, c, l int }
	type neighbor struct{ to, rid, cap int }
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	const mx int = 2e5
	isP := [mx + 1]bool{}
	for i := range isP {
		isP[i] = true
	}
	for i := 2; i <= mx; i++ {
		if isP[i] {
			for j := 2 * i; j <= mx; j += i {
				isP[j] = false
			}
		}
	}

	var n, minP int
	Fscan(in, &n, &minP)
	a := make([]card, n)
	for i := range a {
		Fscan(in, &a[i].p, &a[i].c, &a[i].l)
	}
	ans := sort.Search(n+1, func(upL int) bool {
		g := make([][]neighbor, n+3)
		addEdge := func(from, to, cap int) {
			g[from] = append(g[from], neighbor{to, len(g[to]), cap})
			g[to] = append(g[to], neighbor{from, len(g[from]) - 1, 0})
		}
		st, end, sumP, mxI := n+1, n+2, 0, -1
		for i, c := range a {
			if c.l <= upL && c.c == 1 && (mxI == -1 || c.p > a[mxI].p) {
				mxI = i
			}
		}
		// 源连奇数，偶数连汇
		// 奇数连偶数，特别地，所有 1 中仅选择一个 p 最大的连偶数
		for i, c := range a {
			if c.l > upL || c.c == 1 && i != mxI {
				continue
			}
			sumP += c.p
			if c.c&1 > 0 {
				addEdge(st, i, c.p)
				for j, d := range a {
					if d.l <= upL && d.c&1 == 0 && isP[c.c+d.c] {
						addEdge(i, j, inf)
					}
				}
			} else {
				addEdge(i, end, c.p)
			}
		}

		d := make([]int, n+3)
		bfs := func() bool {
			for i := range d {
				d[i] = -1
			}
			d[st] = 0
			q := []int{st}
			for len(q) > 0 {
				v := q[0]
				q = q[1:]
				for _, e := range g[v] {
					if w := e.to; e.cap > 0 && d[w] < 0 {
						d[w] = d[v] + 1
						q = append(q, w)
					}
				}
			}
			return d[end] >= 0
		}
		var iter []int
		var dfs func(int, int) int
		dfs = func(v, minF int) int {
			if v == end {
				return minF
			}
			for ; iter[v] < len(g[v]); iter[v]++ {
				e := &g[v][iter[v]]
				if w := e.to; e.cap > 0 && d[w] > d[v] {
					if f := dfs(w, min(minF, e.cap)); f > 0 {
						e.cap -= f
						g[w][e.rid].cap += f
						return f
					}
				}
			}
			return 0
		}
		for bfs() {
			iter = make([]int, n+3)
			for {
				if f := dfs(st, inf); f > 0 {
					sumP -= f // 减去最小割
				} else {
					break
				}
			}
		}
		return sumP >= minP
	})
	if ans > n {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF808F(os.Stdin, os.Stdout) }
