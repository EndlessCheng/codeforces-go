package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf123E(in io.Reader, out io.Writer) {
	var n, totA, totB, ans int
	Fscan(in, &n)
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	pa := make([]int, n)
	pb := make([]int, n)
	for i := range pa {
		Fscan(in, &pa[i], &pb[i])
		totA += pa[i]
		totB += pb[i]
	}

	var dfs func(int, int) (int, int)
	dfs = func(v, fa int) (int, int) {
		sumA := pa[v]
		size := 1
		for _, w := range g[v] {
			if w != fa {
				sa, sz := dfs(w, v)
				ans += pb[v] * sa * sz
				sumA += sa
				size += sz
			}
		}
		ans += pb[v] * (totA - sumA) * (n - size)
		return sumA, size
	}
	dfs(0, -1)
	Fprintf(out, "%.9f", float64(ans)/float64(totA*totB))
}

//func main() { cf123E(bufio.NewReader(os.Stdin), os.Stdout) }
