package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"os"
)

type pair struct{ dis, x, y, dir int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dis < h[j].dis }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp) push(v pair)        { heap.Push(h, v) }
func (h *hp) pop() pair          { return heap.Pop(h).(pair) }
func (h hp) empty() bool         { return len(h) == 0 }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type vec struct{ x, y int }
	dir4 := [...]vec{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var n, m, k, x, y, tx, ty, newD int
	Fscan(in, &n, &m, &k, &x, &y, &tx, &ty)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}

	const inf int = 1e18
	dist := make([][][4]int, n)
	for i := range dist {
		dist[i] = make([][4]int, m)
		for j := range dist[i] {
			for k := range dist[i][j] {
				dist[i][j][k] = inf
			}
		}
	}
	dist[x-1][y-1] = [4]int{}
	h := hp{{0, x - 1, y - 1, 0}}
	for !h.empty() {
		p := h.pop()
		d, x, y, dir := p.dis, p.x, p.y, p.dir
		if dist[x][y][dir] < d {
			continue
		}
		for i, di := range dir4 {
			if xx, yy := x+di.x, y+di.y; 0 <= xx && xx < n && 0 <= yy && yy < m && g[xx][yy] == '.' {
				if i == dir {
					newD = d + 1
				} else if d == 0 {
					newD = 1
				} else {
					newD = ((d-1)/k+1)*k + 1
				}
				if newD < dist[xx][yy][i] {
					dist[xx][yy][i] = newD
					h.push(pair{newD, xx, yy, i})
				}
			}
		}
	}
	ans := inf
	for _, d := range dist[tx-1][ty-1] {
		if d < ans {
			ans = d
		}
	}
	if ans < inf {
		ans = (ans-1)/k + 1
	} else {
		ans = -1
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
