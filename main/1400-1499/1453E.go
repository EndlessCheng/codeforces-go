package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1453E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		ans := 0
		var f func(v, fa int) int
		f = func(v, fa int) int {
			mi, mx := -1, 0
			for _, w := range g[v] {
				if w != fa {
					k := f(w, v)
					if mi < 0 || k < mi {
						mi = k
					}
					mx = max(mx, k)
				}
			}
			if len(g[v]) > 2 {
				ans = max(ans, mx+2) // 儿子之间的移动
			}
			return mi + 1 // 从最小的儿子返回
		}
		mx, mx2 := 0, 0
		for _, w := range g[0] {
			k := f(w, 0)
			if k > mx {
				mx, mx2 = k, mx
			} else if k > mx2 {
				mx2 = k
			}
		}
		ans = max(ans, mx+1) // 特殊处理根节点：从最深的儿子返回根节点
		if len(g[0]) > 1 {
			ans = max(ans, mx2+2) // 从非最深的儿子到其他儿子的移动
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1453E(os.Stdin, os.Stdout) }
