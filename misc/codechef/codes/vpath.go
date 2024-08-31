package main

import (
	"bufio"
	. "fmt"
	"io"
)

func vPath(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		const mod = 1_000_000_007
		ans := n
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			s := 0
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				res := dfs(w, v)
				ans = (ans + res*s*2) % mod
				s = (s + res) % mod
			}
			ans = (ans + s) % mod
			return (s*2 + 1) % mod
		}
		dfs(0, -1)
		Fprintln(out, ans)
	}
}

//func main() { vPath(bufio.NewReader(os.Stdin), os.Stdout) }
