package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	size := make([]int, n) // 这是个数据收集者，在各个子树中复用
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		c := a[v] - 1
		old := size[c]
		sz := 1
		for _, w := range g[v] {
			if w != fa {
				size[c] = 0
				szW := dfs(w, v)
				m := szW - size[c]
				ans[c] += m * (m + 1) / 2
				sz += szW
			}
		}
		size[c] = old + sz
		return sz
	}
	dfs(0, -1)
	for i, s := range ans {
		m := n - size[i]
		Fprintln(out, n*(n+1)/2-s-m*(m+1)/2)
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
