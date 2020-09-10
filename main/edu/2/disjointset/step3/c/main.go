package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q int
	Fscan(in, &n, &q)
	if q == 0 {
		return
	}
	type edge struct{ v, w int }
	type query struct {
		edge
		t int
	}
	qs := make([]query, q)
	addTime := map[edge]int{}
	for i := range qs {
		var op string
		if Fscan(in, &op); op[0] == '?' {
			qs[i].t = -1
			continue
		}
		var e edge
		Fscan(in, &e.v, &e.w)
		if e.v > e.w {
			e.v, e.w = e.w, e.v
		}
		qs[i].edge = e
		if op[0] == '+' {
			addTime[e] = i
		} else {
			t := addTime[e]
			delete(addTime, e)
			qs[i].t = t
			qs[t].t = i
		}
	}
	for e, t := range addTime {
		qs[t].t = len(qs)
		qs = append(qs, query{e, t})
	}

	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
	}
	undo := []int{}
	cc := n
	find := func(x int) int {
		for x != fa[x] {
			x = fa[x]
		}
		return x
	}
	merge := func(x, y int) {
		if x, y = find(x), find(y); x != y {
			if sz[x] > sz[y] {
				x, y = y, x
			}
			fa[x] = y
			sz[y] += sz[x]
			undo = append(undo, x)
			cc--
		}
	}
	rollback := func(tar int) {
		for len(undo) > tar {
			x := undo[len(undo)-1]
			undo = undo[:len(undo)-1]
			sz[fa[x]] -= sz[x]
			fa[x] = x
			cc++
		}
	}
	var f func(l, r int)
	f = func(l, r int) {
		if l+1 == r {
			if qs[l].t < 0 {
				Fprintln(out, cc)
			}
			return
		}
		mid := (l + r) >> 1
		checkPoint := len(undo)
		for _, q := range qs[mid:r] {
			if q.t < l {
				merge(q.v, q.w)
			}
		}
		f(l, mid)
		rollback(checkPoint)
		for _, q := range qs[l:mid] {
			if q.t >= r {
				merge(q.v, q.w)
			}
		}
		f(mid, r)
		rollback(checkPoint)
	}
	f(0, len(qs))
}

func main() { run(os.Stdin, os.Stdout) }
