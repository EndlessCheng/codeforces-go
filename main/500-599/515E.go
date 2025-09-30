package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF515E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, h, l, r int
	Fscan(in, &n, &q)
	d := make([]int, n)
	for i := range d {
		Fscan(in, &d[i])
	}
	m := n * 2

	type pair struct{ l, r, h, d int }
	op := func(a, b pair) pair {
		return pair{max(a.l+b.d, b.l), max(a.r, a.d+b.r), max(a.l+b.r, max(a.h, b.h)), a.d + b.d}
	}

	const mx = 18
	st := make([][mx]pair, m)
	for i, d := range d {
		Fscan(in, &h)
		st[i][0] = pair{h*2 + d, h * 2, 0, d}
		st[n+i][0] = st[i][0]
	}
	for j := 1; 1<<j <= m; j++ {
		for i := 0; i+1<<j <= m; i++ {
			st[i][j] = op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	query := func(l, r int) (res pair) {
		for l < r {
			k := bits.Len(uint(r-l)) - 1
			res = op(res, st[l][k])
			l += 1 << k
		}
		return
	}
	for range q {
		Fscan(in, &l, &r)
		l--
		if l < r {
			Fprintln(out, query(r, n+l).h)
		} else {
			Fprintln(out, query(r, l).h)
		}
	}
}

//func main() { CF515E(bufio.NewReader(os.Stdin), os.Stdout) }
