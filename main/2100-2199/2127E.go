package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2127E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
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

		ans := 0
		var dfs func(int, int) map[int]bool
		dfs = func(v, fa int) map[int]bool {
			has := map[int]bool{}
			common := 0
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				sub := dfs(w, v)
				if len(sub) > len(has) {
					has, sub = sub, has
				}
				for c := range sub {
					if !has[c] {
						has[c] = true
					} else if common == 0 {
						common = c
					} else {
						common = -1
					}
				}
			}
			if common < 0 {
				ans += a[v]
				if c[v] == 0 {
					for c[v] = range has {
						break
					}
				}
			} else if common > 0 {
				if c[v] > 0 && c[v] != common {
					ans += a[v]
				} else {
					c[v] = common
				}
			}
			if c[v] > 0 {
				has[c[v]] = true
			}
			return has
		}
		dfs(0, -1)

		var spread func(int, int, int)
		spread = func(v, fa, col int) {
			if c[v] == 0 {
				c[v] = col
			}
			for _, w := range g[v] {
				if w != fa {
					spread(w, v, c[v])
				}
			}
		}
		spread(0, -1, 1)

		Fprintln(out, ans)
		for _, v := range c {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2127E(bufio.NewReader(os.Stdin), os.Stdout) }
