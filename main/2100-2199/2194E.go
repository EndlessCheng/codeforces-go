package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2194E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const inf int = 1e18
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([][]int, n)
		for i := range a {
			a[i] = make([]int, m)
			for j := range a[i] {
				Fscan(in, &a[i][j])
			}
		}

		g := make([][]int, n+1)
		for i := range g {
			g[i] = make([]int, m+1)
		}
		for j := range m - 1 {
			g[n][j] = -inf
		}
		for i := n - 1; i >= 0; i-- {
			g[i][m] = -inf
			for j := m - 1; j >= 0; j-- {
				g[i][j] = max(g[i][j+1], g[i+1][j]) + a[i][j]
			}
		}

		f := make([]int, m+1)
		for j := range f {
			f[j] = -inf
		}
		f[1] = 0
		suf := make([]int, m+1)
		suf[m] = -inf
		ans := inf
		for i, row := range a {
			for j := m - 1; j >= 0; j-- {
				suf[j] = max(suf[j+1], f[j+1]+g[i][j])
			}
			for j, x := range row {
				f[j+1] = max(f[j+1], f[j]) + x
			}
			pre := -inf
			for j, x := range row {
				ans = min(ans, max(pre, f[j+1]+g[i][j]-x*3, suf[j+1]))
				pre = max(pre, f[j+1]+g[i+1][j])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2194E(bufio.NewReader(os.Stdin), os.Stdout) }
