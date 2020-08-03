package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type pair struct{ l, r, i int }

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	qs := make([]pair, q)
	for i := range qs {
		Fscan(in, &qs[i].l, &qs[i].r)
		qs[i].i = i
	}
	sort.Slice(qs, func(i, j int) bool { return qs[i].r < qs[j].r })

	tree := make([]int, n+1)
	add := func(i int, val int) {
		for ; i <= n; i += i & -i {
			tree[i] += val
		}
	}
	sum := func(i int) (res int) {
		for ; i > 0; i &= i - 1 {
			res += tree[i]
		}
		return
	}
	query := func(l, r int) int { return sum(r) - sum(l-1) }

	ans := make([]int, q)
	posR := make([]int, n+1)
	i := 1
	for _, q := range qs {
		for ; i <= q.r; i++ {
			if p := posR[a[i]]; p > 0 {
				add(p, -1)
			}
			add(i, 1)
			posR[a[i]] = i
		}
		ans[q.i] = query(q.l, q.r)
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

func main() { run(os.Stdin, os.Stdout) }
