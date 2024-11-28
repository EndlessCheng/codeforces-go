package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1324F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	color := make([]int, n)
	for i := range color {
		Fscan(in, &color[i])
	}
	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := make([]int, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		res := color[v]*2 - 1
		for _, w := range g[v] {
			if w != fa {
				res += dfs(w, v)
			}
		}
		ans[v] = res
		return max(res, 0)
	}
	dfs(0, -1)

	var reroot func(int, int)
	reroot = func(v, fa int) {
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			if ans[w] < 0 {
				ans[w] += max(ans[v], 0)
			} else {
				ans[w] = max(ans[w], ans[v])
			}
			reroot(w, v)
		}
	}
	reroot(0, -1)

	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF1324F(bufio.NewReader(os.Stdin), os.Stdout) }
