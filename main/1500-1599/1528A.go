package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1528A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	abs := func(x int) int64 {
		if x < 0 {
			return int64(-x)
		}
		return int64(x)
	}

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]struct{ l, r int }, n)
		for i := range a {
			Fscan(in, &a[i].l, &a[i].r)
		}
		type nb struct{ to, wt int }
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		var f func(v, fa int) (sl, sr int64)
		f = func(v, fa int) (sl, sr int64) {
			for _, w := range g[v] {
				if w != fa {
					p, q := f(w, v)
					sl += max(abs(a[v].l-a[w].l)+p, abs(a[v].l-a[w].r)+q)
					sr += max(abs(a[v].r-a[w].l)+p, abs(a[v].r-a[w].r)+q)
				}
			}
			return
		}
		Fprintln(out, max(f(0, -1)))
	}
}

//func main() { CF1528A(os.Stdin, os.Stdout) }
