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

	var dfs func(int, []int) ([]int, int)
	dfs = func(v int, pre []int) ([]int, int) {
		size := 1
		t := pre
		for _, w := range g[v] {
			f, sz := dfs(w, t)
			t = f
			size += sz
		}
		f := slices.Clone(pre)
		p := a[v]
		for j := W; j >= p.w; j-- {
			f[j] = max(f[j], t[j-p.w]+p.v)
		}
		return f, size
	}
	f, _ := dfs(0, make([]int, W+1))
	Fprint(out, f[W])
}

//func main() { p1064(bufio.NewReader(os.Stdin), os.Stdout) }
