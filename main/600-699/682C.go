package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF682C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, wt, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	type nb struct{ to, wt int }
	g := make([][]nb, n)
	for w := 1; w < n; w++ {
		Fscan(in, &v, &wt)
		g[v-1] = append(g[v-1], nb{w, wt})
	}

	var f func(int, int64, bool)
	f = func(v int, mx int64, del bool) {
		if mx > int64(a[v]) {
			del = true
		}
		if del {
			ans++
		}
		for _, e := range g[v] {
			m := mx + int64(e.wt)
			if m < 0 {
				m = 0
			}
			f(e.to, m, del)
		}
	}
	f(0, 0, false)
	Fprint(out, ans)
}

//func main() { CF682C(os.Stdin, os.Stdout) }
