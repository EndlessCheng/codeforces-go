package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p1273(in io.Reader, out io.Writer) {
	var n, m, k int
	Fscan(in, &n, &m)
	g := make([][]int, n-m)
	a := make([]int, n)
	for i := range g {
		Fscan(in, &k)
		g[i] = make([]int, k)
		for j := range g[i] {
			var v, w int
			Fscan(in, &v, &w)
			v--
			g[i][j] = v
			a[v] = -w
		}
	}
	for i := n - m; i < n; i++ {
		var v int
		Fscan(in, &v)
		a[i] += v
	}

	var dfs func(int) ([]int, int)
	dfs = func(v int) ([]int, int) {
		if v >= n-m {
			return []int{0, a[v]}, 1
		}
		f := make([]int, m+1)
		for i := 1; i <= m; i++ {
			f[i] = -1e9
		}
		son := 0
		for _, w := range g[v] {
			fw, s := dfs(w)
			for i := son; i >= 0; i-- {
				for j, wCj := range fw {
					f[i+j] = max(f[i+j], f[i]+wCj)
				}
			}
			son += s
		}
		for i := 1; i <= son; i++ {
			f[i] += a[v]
		}
		return f[:son+1], son
	}
	f, _ := dfs(0)
	for i := m; ; i-- {
		if f[i] >= 0 {
			Fprint(out, i)
			return
		}
	}
}

//func main() { p1273(bufio.NewReader(os.Stdin), os.Stdout) }
