package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF707D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type query struct{ t, r, c int }
	var n, m, q, tot int
	Fscan(in, &n, &m, &q)
	qs := make([]query, q+1)
	g := make([][]int, q+1)
	for i := 1; i <= q; i++ {
		if Fscan(in, &qs[i].t, &qs[i].r); qs[i].t < 4 {
			qs[i].r--
			if qs[i].t < 3 {
				Fscan(in, &qs[i].c)
				qs[i].c--
			}
			g[i-1] = append(g[i-1], i)
		} else {
			g[qs[i].r] = append(g[qs[i].r], i)
		}
	}

	ans := make([]int, q+1)
	has := make([][]bool, n)
	for i := range has {
		has[i] = make([]bool, m)
	}
	inv := make([]bool, n)
	cnt := make([]int, n)
	update := func(q *query) {
		switch r, c := q.r, q.c; q.t {
		case 1:
			if has[r][c] == inv[r] {
				has[r][c] = !inv[r]
				tot++
				cnt[r]++
				q.t = 2 // 方便撤销
			} else {
				q.t = 0
			}
		case 2:
			if has[r][c] != inv[r] {
				has[r][c] = inv[r]
				tot--
				cnt[r]--
				q.t = 1 // 方便撤销
			} else {
				q.t = 0
			}
		case 3:
			inv[r] = !inv[r]
			tot += m - cnt[r]*2
			cnt[r] = m - cnt[r]
		}
	}
	var f func(int)
	f = func(v int) {
		for _, w := range g[v] {
			update(&qs[w])
			ans[w] = tot
			f(w)
			update(&qs[w]) // 撤销
		}
	}
	f(0)
	for _, v := range ans[1:] {
		Fprintln(out, v)
	}
}

//func main() { CF707D(os.Stdin, os.Stdout) }
