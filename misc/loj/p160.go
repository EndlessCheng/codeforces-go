package main

import (
	. "fmt"
	"io"
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

	f := make([][]int, 1, n+1)
	f[0] = make([]int, W+1)
	var dfs func(int) int
	dfs = func(v int) int {
		size := 1
		for _, w := range g[v] {
			size += dfs(w)
		}
		t := f[len(f)-size]
		cur := append(t[:0:0], t...)
		p := a[v]
		t = f[len(f)-1]
		for j := W; j >= p.w; j-- {
			cur[j] = max(cur[j], t[j-p.w]+p.v)
		}
		f = append(f, cur)
		return size
	}
	dfs(0)
	Fprint(out, f[n][W])
}

//func main() { p160(bufio.NewReader(os.Stdin), os.Stdout) }
