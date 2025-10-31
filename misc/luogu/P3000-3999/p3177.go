package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func p3177(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	m = min(m, n-m)
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for i := 1; i < n; i++ {
		var v, w, wt int
		Fscan(in, &v, &w, &wt)
		v--
		w--
		g[v] = append(g[v], nb{w, wt})
		g[w] = append(g[w], nb{v, wt})
	}

	var dfs func(int, int) ([]int, int)
	dfs = func(v, fa int) ([]int, int) {
		f := make([]int, m+1)
		size := 1
		for _, e := range g[v] {
			w := e.to
			if w == fa {
				continue
			}
			fw, sz := dfs(w, v)
			size += sz
			// 枚举在 v 子树内选 j 个黑点
			for j := min(size, m); j >= 0; j-- {
				// 枚举在 w 子树内选 k 个黑点（其余子树选 j-k 个黑点，那么有 j-k <= size-sz）
				for k := max(j-size+sz, 0); k <= min(j, sz); k++ {
					f[j] = max(f[j], f[j-k]+fw[k]+(k*(m-k)+(sz-k)*(n-m-sz+k))*e.wt)
				}
			}
		}
		return f, size
	}
	f, _ := dfs(0, -1)
	Fprint(out, f[m])
}

//func main() { p3177(bufio.NewReader(os.Stdin), os.Stdout) }
