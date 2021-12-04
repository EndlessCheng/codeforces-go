package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF743D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	max := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	const inf int64 = 1e18

	var n, v, w int
	Fscan(in, &n)
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

	ans := -inf
	var f func(int, int) (int64, int64)
	f = func(v, fa int) (int64, int64) {
		sum, fi, se := int64(a[v]), -inf*2, -inf*2
		for _, w := range g[v] {
			if w != fa {
				s, mx := f(w, v)
				sum += s
				if mx > fi {
					fi, se = mx, fi
				} else if mx > se {
					se = mx
				}
			}
		}
		ans = max(ans, fi+se)
		return sum, max(sum, fi)
	}
	f(0, -1)
	if ans == -inf {
		Fprint(out, "Impossible")
	} else {
		Fprint(out, ans)
	}
}

//func main() { CF743D(os.Stdin, os.Stdout) }
