package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1042F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v, w, rt, ans int
	Fscan(in, &n, &k)
	g := make([][]int, n+1)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}
	for i, vs := range g {
		if len(vs) > 1 {
			rt = i
			break
		}
	}
	var dfs func(int, int) int
	dfs = func(v, fa int) (mx int) {
		if len(g[v]) == 1 {
			return 1
		}
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			d := dfs(w, v)
			if mx+d <= k {
				mx = max(mx, d)
			} else {
				ans++
				mx = min(mx, d)
			}
		}
		if mx > 0 {
			mx++
		}
		return
	}
	d := dfs(rt, -1)
	if d > 0 {
		ans++
	}
	Fprint(out, ans)
}

//func main() { cf1042F(os.Stdin, os.Stdout) }
