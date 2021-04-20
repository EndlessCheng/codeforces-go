package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1092F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	a := make([]int64, n)
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

	ans := int64(0)
	s := make([]int64, n)
	var f func(v, fa, d int) int64
	f = func(v, fa, d int) int64 {
		ans += a[v] * int64(d)
		s[v] = a[v]
		for _, w := range g[v] {
			if w != fa {
				s[v] += f(w, v, d+1)
			}
		}
		return s[v]
	}
	f(0, -1, 0)
	var f2 func(int, int, int64)
	f2 = func(v, fa int, sum int64) {
		if sum > ans {
			ans = sum
		}
		for _, w := range g[v] {
			if w != fa {
				f2(w, v, sum+s[0]-s[w]*2)
			}
		}
	}
	f2(0, -1, ans)
	Fprint(out, ans)
}

//func main() { CF1092F(os.Stdin, os.Stdout) }
