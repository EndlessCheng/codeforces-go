package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF685B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m, v int
	Fscan(in, &n, &m)
	pa := make([]int, n+1)
	g := make([][]int, n+1)
	for w := 2; w <= n; w++ {
		Fscan(in, &pa[w])
		g[pa[w]] = append(g[pa[w]], w)
	}
	ct := make([]int, n+1)
	sz := make([]int, n+1)
	var f func(int)
	f = func(v int) {
		ct[v] = v
		sz[v] = 1
		mx := 0
		for _, w := range g[v] {
			f(w)
			sz[v] += sz[w]
			if sz[w] > sz[mx] {
				mx = w
			}
		}
		if sz[mx]*2 > sz[v] {
			x := ct[mx]
			for sz[x]*2 < sz[v] {
				x = pa[x]
			}
			ct[v] = x
		}
	}
	f(1)
	for ; m > 0; m-- {
		Fscan(in, &v)
		Fprintln(out, ct[v])
	}
}

//func main() { CF685B(os.Stdin, os.Stdout) }
