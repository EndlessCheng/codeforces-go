package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1777D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod int64 = 1e9 + 7
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		p2 := int64(1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
			p2 = p2 * 2 % mod
		}

		ans := int64(0)
		var f func(v, fa int) int
		f = func(v, fa int) (dep int) {
			for _, w := range g[v] {
				if w != fa {
					dep = max(dep, f(w, v))
				}
			}
			dep++
			ans += int64(dep)
			return
		}
		f(0, -1)
		Fprintln(out, ans%mod*p2%mod)
	}
}

//func main() { CF1777D(os.Stdin, os.Stdout) }
