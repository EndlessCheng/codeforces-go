package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, q, l, r int
	Fscan(in, &n, &m)
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
	type edge struct{ v, w int }
	merge := func(e edge) {
		if x, y := find(e.v), find(e.w); x != y {
			if sz[x] > sz[y] {
				x, y = y, x
			}
			fa[x] = y
			sz[y] += sz[x]
			undo = append(undo, x)
			cc--
		}
	}
	rollback := func(k int) {
		for len(undo) > k {
			x := undo[len(undo)-1]
			undo = undo[:len(undo)-1]
			sz[fa[x]] -= sz[x]
			fa[x] = x
			cc++
		}
	}

	es := make([]edge, m)
	for i := range es {
		Fscan(in, &es[i].v, &es[i].w)
	}

	Fscan(in, &q)
	ans := make([]interface{}, q)
	type query struct{ l, r, i int }
	blockSize := int(math.Round(math.Sqrt(float64(m))))
	qs := make([][]query, (m-1)/blockSize+1)
	for i := 0; i < q; i++ {
		Fscan(in, &l, &r)
		l--
		j := l / blockSize
		qs[j] = append(qs[j], query{l, r, i})
	}
	for i, b := range qs {
		r := (i + 1) * blockSize
		maxR := r
		sort.Slice(b, func(i, j int) bool { return b[i].r < b[j].r })
		for _, q := range b {
			if q.r <= r {
				for _, e := range es[q.l:q.r] {
					merge(e)
				}
				ans[q.i] = cc
				rollback(0)
			} else {
				for _, e := range es[maxR:q.r] {
					merge(e)
				}
				maxR = q.r
				checkPoint := len(undo)
				for _, e := range es[q.l:r] {
					merge(e)
				}
				ans[q.i] = cc
				rollback(checkPoint)
			}
		}
		rollback(0)
	}
	Fprintln(out, ans...)
}

func main() { run(os.Stdin, os.Stdout) }
