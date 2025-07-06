package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
type tuple struct {
	w, v int
	f    [2001]int
}

type stack []tuple

func (st stack) res(w int) int {
	return st[len(st)-1].f[w]
}

func (st *stack) push(w, v int) {
	f := (*st)[len(*st)-1].f
	for i := len(f) - 1; i >= w; i-- {
		f[i] = max(f[i], f[i-w]+v)
	}
	*st = append(*st, tuple{w, v, f})
}

func (st *stack) pop() (w, v int) {
	n := len(*st) - 1
	w, v = (*st)[n].w, (*st)[n].v
	*st = (*st)[:n]
	return
}

func (st stack) empty() bool {
	return len(st) == 1
}

type deque struct{ l, r stack }

func (q *deque) rebalance() {
	if q.r.empty() {
		q.l, q.r = q.r, q.l
		defer func() { q.l, q.r = q.r, q.l }()
	}
	m := len(q.r) / 2
	for i := m; i > 0; i-- {
		q.l.push(q.r[i].w, q.r[i].v)
	}
	t := q.r[m+1:]
	q.r = q.r[:1]
	for _, p := range t {
		q.r.push(p.w, p.v)
	}
}

func (q deque) res(w int) (mx int) {
	for i := range w + 1 {
		mx = max(mx, q.l.res(i)+q.r.res(w-i))
	}
	return
}

func (q *deque) pushFront(w, v int) {
	q.l.push(w, v)
}

func (q *deque) pushBack(w, v int) {
	q.r.push(w, v)
}

func (q *deque) popFront() (w, v int) {
	if q.l.empty() {
		q.rebalance()
	}
	return q.l.pop()
}

func (q *deque) popBack() {
	if q.r.empty() {
		q.rebalance()
	}
	q.r.pop()
}

func cf2026F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, nodeId, op, x, p, t int
	Fscanln(in, &q)
	type edge struct{ to, op, p, t, i int }
	g := make([][]edge, q+1)
	pos := make([]int, q+1)
	store := 1

	for i := range q {
		Fscanln(in, &op, &x, &p, &t)
		nodeId++
		v := pos[x]
		g[v] = append(g[v], edge{nodeId, op, p, t, i})
		if op == 1 {
			store++
			pos[store] = nodeId
		} else {
			pos[x] = nodeId
		}
	}

	ans := make([]int, q)
	dq := deque{stack{{}}, stack{{}}}
	var dfs func(int)
	dfs = func(v int) {
		for _, e := range g[v] {
			if e.op == 2 {
				dq.pushBack(e.p, e.t)
			} else if e.op == 3 {
				e.p, e.t = dq.popFront()
			} else if e.op == 4 {
				ans[e.i] = dq.res(e.p) + 1
			}
			dfs(e.to)
			if e.op == 2 {
				dq.popBack()
			} else if e.op == 3 {
				dq.pushFront(e.p, e.t)
			}
		}
	}
	dfs(0)

	for _, v := range ans {
		if v > 0 {
			Fprintln(out, v-1)
		}
	}
}

//func main() { cf2026F(bufio.NewReader(os.Stdin), os.Stdout) }
