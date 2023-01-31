package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, wt, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		mx := a[v]
		ans = max(ans, mx)
		for _, e := range g[v] {
			if w := e.to; w != fa {
				if m := dfs(w, v) - e.wt; m >= 0 {
					ans = max(ans, mx+m) // 拼接来自两棵不同子树的路径
					mx = max(mx, m+a[v]) // 更新子树路径的最大值
				}
			}
		}
		return mx
	}
	dfs(0, -1)
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
