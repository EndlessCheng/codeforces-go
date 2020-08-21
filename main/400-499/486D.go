package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF486D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var d, n, v, w, rt, min int
	Fscan(in, &d, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	const mod int64 = 1e9 + 7
	var f func(v, fa int) int64
	f = func(v, fa int) int64 {
		s := int64(1)
		for _, w := range g[v] {
			if w != fa && (w > rt && a[w] == min || min < a[w] && a[w] <= min+d) {
				s = s * (f(w, v) + 1) % mod
			}
		}
		return s
	}
	ans := int64(0)
	for rt, min = range a {
		ans += f(rt, -1)
	}
	Fprint(out, ans%mod)
}

//func main() { CF486D(os.Stdin, os.Stdout) }
