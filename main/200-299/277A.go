package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF277A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, k, v, ans int
	Fscan(in, &n, &m)
	g := make([][]int, n+m)
	ok := 0
	for i := 0; i < n; i++ {
		for Fscan(in, &k); k > 0; k-- {
			Fscan(in, &v)
			v += n - 1
			g[i] = append(g[i], v)
			g[v] = append(g[v], i)
			ok = 1
		}
	}

	vis := make([]bool, len(g))
	var f func(int)
	f = func(v int) {
		vis[v] = true
		for _, w := range g[v] {
			if !vis[w] {
				f(w)
			}
		}
	}
	for i, b := range vis[:n] {
		if !b {
			ans++
			f(i)
		}
	}
	Fprint(out, ans-ok)
}

//func main() { CF277A(os.Stdin, os.Stdout) }
