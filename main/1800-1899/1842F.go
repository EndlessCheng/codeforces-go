package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf1842F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
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

	dep := make([]int, n)
	var dfs func(int, int)
	dfs = func(v, fa int) {
		for _, w := range g[v] {
			if w != fa {
				dep[w] = dep[v] + 1
				dfs(w, v)
			}
		}
	}

	ans := make([]int, n+1)
	for i := range n {
		dep[i] = 0
		dfs(i, -1)

		slices.Sort(dep)
		s := 0
		for j, d := range dep {
			j++
			s += d
			ans[j] = max(ans[j], (n-1)*j-s*2)
		}
	}

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { cf1842F(bufio.NewReader(os.Stdin), os.Stdout) }
