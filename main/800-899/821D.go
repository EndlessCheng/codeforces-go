package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type deque21 struct{ l, r []int }

func (q deque21) empty() bool  { return len(q.l) == 0 && len(q.r) == 0 }
func (q *deque21) pushL(v int) { q.l = append(q.l, v) }
func (q *deque21) pushR(v int) { q.r = append(q.r, v) }
func (q *deque21) popL() (v int) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

func CF821D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	const inf int = 1e9
	type pair struct{ x, y int }
	var n, m, k, x, y int
	Fscan(in, &n, &m, &k)
	id := map[pair]int{}
	for i := 0; i < k; i++ {
		Fscan(in, &x, &y)
		id[pair{x - 1, y - 1}] = i
	}

	type nb struct{ to, wt int }
	g := make([][]nb, k+n+m)
	add := func(v, w int) {
		g[v] = append(g[v], nb{w, 1})
		g[w] = append(g[w], nb{v, 0})
	}
	for p, i := range id {
		x, y := p.x, p.y
		if x > 0 {
			if j, has := id[pair{x - 1, y}]; has {
				g[i] = append(g[i], nb{j, 0})
				g[j] = append(g[j], nb{i, 0})
			}
		}
		if y > 0 {
			if j, has := id[pair{x, y - 1}]; has {
				g[i] = append(g[i], nb{j, 0})
				g[j] = append(g[j], nb{i, 0})
			}
		}
		add(i, k+x)
		if x > 0 {
			add(i, k+x-1)
		}
		if x+1 < n {
			add(i, k+x+1)
		}
		add(i, k+n+y)
		if y > 0 {
			add(i, k+n+y-1)
		}
		if y+1 < m {
			add(i, k+n+y+1)
		}
	}

	st := id[pair{}]
	dis := make([]int, len(g))
	for i := range dis {
		dis[i] = inf
	}
	dis[st] = 0
	q := &deque21{}
	q.pushL(st)
	for !q.empty() {
		v := q.popL()
		for _, e := range g[v] {
			w, d := e.to, e.wt
			if newD := dis[v] + d; newD < dis[w] {
				dis[w] = newD
				if d == 0 {
					q.pushL(w)
				} else {
					q.pushR(w)
				}
			}
		}
	}
	ans := inf
	if v, has := id[pair{n - 1, m - 1}]; has {
		ans = dis[v]
	}
	ans = min(ans, min(dis[k+n-1], dis[k+n+m-1]))
	if ans == inf {
		ans = -1
	}
	Fprint(out, ans)
}

//func main() { CF821D(os.Stdin, os.Stdout) }
