package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1406C(_r io.Reader, _w io.Writer) {
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
		g := make([][]int, n+1)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		ct := []int{}
		min := int(1e9)
		var f func(v, fa int) int
		f = func(v, fa int) int {
			sz, maxSz := 1, 0
			for _, w := range g[v] {
				if w != fa {
					s := f(w, v)
					sz += s
					maxSz = max(maxSz, s)
				}
			}
			maxSz = max(maxSz, n-sz)
			if maxSz < min {
				min = maxSz
				ct = []int{v}
			} else if maxSz == min {
				ct = append(ct, v)
			}
			return sz
		}
		f(1, 0)
		if len(ct) == 1 {
			Fprintln(out, v, w)
			Fprintln(out, v, w)
		} else {
			v, w = ct[0], ct[1]
			u := g[v][0]
			if u == w {
				u = g[v][1]
			}
			Fprintln(out, v, u)
			Fprintln(out, w, u)
		}
	}
}

//func main() { CF1406C(os.Stdin, os.Stdout) }
