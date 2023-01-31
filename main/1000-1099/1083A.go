package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1083A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}

	var n, v, w, wt int
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

	ans := int64(0)
	var f func(int, int) int64
	f = func(v, fa int) int64 {
		mx := int64(a[v])
		ans = max(ans, mx)
		for _, e := range g[v] {
			if w := e.to; w != fa {
				if m := f(w, v) - int64(e.wt); m >= 0 {
					ans = max(ans, mx+m)        // 拼接来自两棵不同子树的路径
					mx = max(mx, m+int64(a[v])) // 更新子树路径的最大值
				}
			}
		}
		return mx
	}
	f(0, -1)
	Fprint(out, ans)
}

//func main() { CF1083A(os.Stdin, os.Stdout) }
