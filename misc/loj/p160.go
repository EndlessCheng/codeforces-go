package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p160(in io.Reader, out io.Writer) {
	var n, W int
	Fscan(in, &n, &W)
	g := make([][]int, n+1)
	for w := 1; w < n+1; w++ {
		var p int
		Fscan(in, &p)
		g[p] = append(g[p], w)
	}
	a := make([]struct{ w, v int }, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].w)
	}
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].v)
	}

	var dfs func(int, []int) []int
	dfs = func(v int, pre []int) []int {
		t := pre
		for _, w := range g[v] {
			t = dfs(w, t)
		}
		f := slices.Clone(pre)
		p := a[v]
		for j := W; j >= p.w; j-- {
			f[j] = max(f[j], t[j-p.w]+p.v)
		}
		return f
	}
	f := dfs(0, make([]int, W+1))
	Fprint(out, f[W])
}

//func main() { p160(bufio.NewReader(os.Stdin), os.Stdout) }
