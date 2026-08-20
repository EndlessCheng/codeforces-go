package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2227H(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n+1)
		d := make([]int, n+2)
		diff := make([]int, n+2)
		r := make([]int, n+2)
		for range n - 1 {
			var v, w int
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
			d[v]++
			d[w]++
		}

		ans, p := 0, 0

		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			p++
			l := p
			s := 0
			if d[v] == 1 {
				s = 1
			}
			r[l] = s

			for _, w := range g[v] {
				if w != fa {
					s ^= dfs(w, v)
				}
			}

			if fa != 0 {
				diff[1] += s
				if ans != 0 {
					diff[l] -= s*2 - 1
					diff[p+1] += s*2 - 1
				}
			}
			return s
		}

		for i := 1; i <= n; i++ {
			if d[i] == 1 {
				ans ^= 1
			}
		}
		dfs(1, 0)

		ans = n
		for i := 0; i <= n; i++ {
			if i > 0 {
				diff[i] += diff[i-1]
			}
			if r[i] != 0 {
				ans = min(ans, diff[i])
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { cf2227H(bufio.NewReader(os.Stdin), os.Stdout) }
