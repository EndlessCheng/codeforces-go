package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1728G(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353
	var n, m, q, base int
	Fscan(in, &base, &n, &m)
	a := make([]int, n+2)
	b := make([]int, n+2)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		b[i] = base
	}

	c := make([]int, m+2)
	for i := 1; i <= m; i++ {
		Fscan(in, &c[i])
	}
	slices.Sort(c[1 : m+1])
	c[0] = -1e9
	c[m+1] = 1e9

	g := make([][]int, m+2)
	for i := range g {
		g[i] = make([]int, m+1)
	}
	for i := 1; i <= m+1; i++ {
		for j := range i {
			g[i][j] = 1
			for k := 1; k <= n; k++ {
				if c[j] <= a[k] && a[k] <= c[i] {
					g[i][j] = g[i][j] * min(a[k]-c[j], c[i]-a[k], b[k]+1) % mod
				}
			}
		}
	}

	f := make([]int, m+2)
	Fscan(in, &q)
	for range q {
		Fscan(in, &a[n+1])
		b[n+1] = base
		f[0] = 1
		for i := 1; i <= m+1; i++ {
			f[i] = 0
			for j := range i {
				way := g[i][j]
				if c[j] <= a[n+1] && a[n+1] <= c[i] {
					way = way * min(a[n+1]-c[j], c[i]-a[n+1], b[n+1]+1) % mod
				}
				f[i] = (f[i] - f[j]*way) % mod
			}
		}
		Fprintln(out, (mod-f[m+1])%mod)
	}
}

//func main() { cf1728G(bufio.NewReader(os.Stdin), os.Stdout) }
