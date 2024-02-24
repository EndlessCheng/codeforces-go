package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1923E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		cs := make([]int, n)
		for i := range cs {
			Fscan(in, &cs[i])
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ans := 0
		cnt := make([]int, n+1)
		var dfs func(int, int)
		dfs = func(v, fa int) {
			c := cs[v]
			tmp := cnt[c]
			ans += tmp
			for _, w := range g[v] {
				if w != fa {
					cnt[c] = 1
					dfs(w, v)
				}
			}
			cnt[c] = tmp + 1
		}
		dfs(0, -1)
		Fprintln(out, ans)
	}
}

//func main() { cf1923E(os.Stdin, os.Stdout) }
