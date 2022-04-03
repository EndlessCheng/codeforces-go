package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(Case int) {
		var n, v, ans int
		Fscan(in, &n)
		a := make([]int, n+1)
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
		}
		g := make([][]int, n+1)
		for w := 1; w <= n; w++ {
			Fscan(in, &v)
			g[v] = append(g[v], w)
		}

		var f func(int) int
		f = func(v int) int {
			if len(g[v]) == 0 {
				ans += a[v]
				return a[v]
			}
			mi := int(1e9)
			for _, w := range g[v] {
				mi = min(mi, f(w))
			}
			if a[v] > mi {
				ans += a[v] - mi
				mi = a[v]
			}
			return mi
		}
		for _, v := range g[0] {
			f(v)
		}
		Fprintln(out, ans)
	}

	var T int
	Fscan(in, &T)
	for Case := 1; Case <= T; Case++ {
		Fprintf(out, "Case #%d: ", Case)
		solve(Case)
	}
}

func main() { run(os.Stdin, os.Stdout) }

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
