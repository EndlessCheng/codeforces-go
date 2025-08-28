package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1101F(in io.Reader, out io.Writer) {
	var n, m, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	type query struct{ r, c, k int }
	g := make([][]query, n)
	for range m {
		var l, r, c, k int
		Fscan(in, &l, &r, &c, &k)
		g[l-1] = append(g[l-1], query{r - 1, c, min(k+1, r-l)})
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for i := range f[0] {
		f[0][i] = 1e9
	}
	for l, qs := range g {
		f[0][l] = 0
		for k := 1; k < n-l; k++ {
			opt := l
			for r := l + 1; r < n; r++ {
				for a[r]-a[opt+1] > f[k-1][opt+1] {
					opt++
				}
				f[k][r] = min(a[r]-a[opt], f[k-1][opt+1])
			}
		}
		for _, q := range qs {
			ans = max(ans, f[q.k][q.r]*q.c)
		}
	}
	Fprint(out, ans)
}

//func main() { cf1101F(bufio.NewReader(os.Stdin), os.Stdout) }
