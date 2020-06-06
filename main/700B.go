package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF700B(_r io.Reader, out io.Writer) {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	in := bufio.NewReader(_r)
	var n, k, v, w int
	Fscan(in, &n, &k)
	k *= 2
	has := make([]bool, n)
	for i := k; i > 0; i-- {
		Fscan(in, &v)
		has[v-1] = true
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := int64(0)
	var f func(v, fa int) int
	f = func(v, fa int) (sz int) {
		if has[v] {
			sz = 1
		}
		for _, w := range g[v] {
			if w != fa {
				s := f(w, v)
				ans += int64(min(s, k-s))
				sz += s
			}
		}
		return
	}
	f(0, -1)
	Fprint(out, ans)
}

//func main() { CF700B(os.Stdin, os.Stdout) }
