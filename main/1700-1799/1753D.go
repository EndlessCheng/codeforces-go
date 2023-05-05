package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1753D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const inf int64 = 1e18
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	match := []struct{ x, y int }{'L': {0, 1}, 'R': {0, -1}, 'U': {1, 0}, 'D': {-1, 0}}
	min := func(a, b int64) int64 {
		if a > b {
			return b
		}
		return a
	}

	var n, m int
	var p, q int64
	Fscan(in, &n, &m, &p, &q)
	g := make([]string, n)
	dis := make([][]int64, n)
	h := hp53{}
	for i := range g {
		Fscan(in, &g[i])
		dis[i] = make([]int64, m)
		for j, b := range g[i] {
			if b == '.' {
				dis[i][j] = 0
				h.push(pd53{i, j, 0})
			} else {
				dis[i][j] = inf
			}
		}
	}
	for len(h) > 0 {
		top := h.pop()
		i, j := top.x, top.y
		if top.dis > dis[i][j] {
			continue
		}
		for _, d := range dir4 {
			x, y := i+d.x, j+d.y
			if 0 <= x && x < n && 0 <= y && y < m && g[x][y] != '.' && g[x][y] != '#' {
				d := match[g[x][y]]
				x += d.x
				y += d.y
				newD := dis[i][j]
				if x != i && y != j {
					newD += p
				} else {
					newD += q
				}
				if newD < dis[x][y] {
					dis[x][y] = newD
					h.push(pd53{x, y, newD})
				}
			}
		}
	}
	ans := inf
	for i, r := range dis {
		for j, d := range r {
			if j < m-1 {
				ans = min(ans, d+r[j+1])
			}
			if i < n-1 {
				ans = min(ans, d+dis[i+1][j])
			}
		}
	}
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF1753D(os.Stdin, os.Stdout) }

type pd53 struct {
	x, y int
	dis  int64
}
type hp53 []pd53

func (h hp53) Len() int              { return len(h) }
func (h hp53) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h hp53) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp53) Push(v interface{})   { *h = append(*h, v.(pd53)) }
func (h *hp53) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *hp53) push(v pd53)          { heap.Push(h, v) }
func (h *hp53) pop() pd53            { return heap.Pop(h).(pd53) }
