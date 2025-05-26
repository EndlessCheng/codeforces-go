package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p1064(in io.Reader, out io.Writer) {
	var W, n, v int
	Fscan(in, &W, &n)
	a := make([]struct{ w, v int }, n+1)
	g := make([][]int, n+1)
	for w := 1; w <= n; w++ {
		Fscan(in, &a[w].w, &a[w].v, &v)
		a[w].v *= a[w].w
		g[v] = append(g[v], w)
	}

	var dfs func(int, []int) []int
	dfs = func(v int, t []int) []int {
		f := slices.Clone(t)
		for _, w := range g[v] {
			t = dfs(w, t)
		}
		p := a[v]
		for j := W; j >= p.w; j-- {
			f[j] = max(f[j], t[j-p.w]+p.v)
		}
		return f
	}
	t := make([]int, W+1)
	Fprint(out, dfs(0, t)[W])
}

//func main() { p1064(bufio.NewReader(os.Stdin), os.Stdout) }
