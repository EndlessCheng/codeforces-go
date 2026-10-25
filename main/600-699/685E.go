package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf685E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, q int
	Fscan(in, &n, &m, &q)
	x := make([]int, m+1)
	y := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &x[i], &y[i])
	}

	type query struct {r, s, t, i int}
	qs := make([][]query, m+1)
	for i := range q {
		var l, r, s, t int
		Fscan(in, &l, &r, &s, &t)
		qs[l] = append(qs[l], query{r, s, t, i})
	}

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
		for j := range f[i] {
			f[i][j] = 1e18
		}
	}
	ans := make([]bool, q)
	for i := m; i > 0; i-- {
		f[x[i]][y[i]] = i
		f[y[i]][x[i]] = i
		for j := 1; j <= n; j++ {
			f[x[i]][j] = min(f[x[i]][j], f[y[i]][j])
			f[y[i]][j] = min(f[x[i]][j], f[y[i]][j])
		}
		for _, q := range qs[i] {
			ans[q.i] = f[q.s][q.t] <= q.r
		}
	}

	for _, b := range ans {
		if b {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() { cf685E(bufio.NewReader(os.Stdin), os.Stdout) }
