package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://space.bilibili.com/206214
func p1064(in io.Reader, out io.Writer) {
	var maxW, n, v int
	Fscan(in, &maxW, &n)
	a := make([]struct{ w, v int }, n+1)
	g := make([][]int, n+1)
	for w := 1; w <= n; w++ {
		Fscan(in, &a[w].w, &a[w].v, &v)
		a[w].v *= a[w].w
		g[v] = append(g[v], w)
	}

	f := make([][]int, 1, n+1)
	f[0] = make([]int, maxW+1)
	var dfs func(int) int
	dfs = func(v int) int {
		size := 1
		for _, w := range g[v] {
			size += dfs(w)
		}
		fv := slices.Clone(f[len(f)-size])
		p := a[v]
		lastF := f[len(f)-1]
		for j := maxW; j >= p.w; j-- {
			fv[j] = max(fv[j], lastF[j-p.w]+p.v)
		}
		f = append(f, fv)
		return size
	}
	dfs(0)
	Fprint(out, f[n][maxW])
}

//func main() { p1064(bufio.NewReader(os.Stdin), os.Stdout) }
