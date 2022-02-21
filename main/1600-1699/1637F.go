package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1637F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w, root, fi, se int
	Fscan(in, &n)
	h := make([]int, n)
	for i := range h {
		Fscan(in, &h[i])
		if h[i] > h[root] {
			root = i
		}
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
	var f func(int, int) int
	f = func(v, fa int) (max int) {
		for _, w := range g[v] {
			if w != fa {
				if mx := f(w, v); mx > max {
					max = mx
				}
			}
		}
		if h[v] > max {
			ans += int64(h[v] - max)
			max = h[v]
		}
		return
	}
	for _, w := range g[root] {
		if mx := f(w, root); mx > fi {
			fi, se = mx, fi
		} else if mx > se {
			se = mx
		}
	}
	Fprint(out, ans+int64(h[root]*2-fi-se))
}

//func main() { CF1637F(os.Stdin, os.Stdout) }
