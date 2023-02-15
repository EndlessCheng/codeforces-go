package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1176E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n+1)
		for ; m > 0; m-- {
			Fscan(in, &v, &w)
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		vis := make([]bool, len(g))
		a := [2][]int{}
		var f func(int, int)
		f = func(v, d int) {
			vis[v] = true
			a[d] = append(a[d], v)
			for _, w := range g[v] {
				if !vis[w] {
					f(w, d^1)
				}
			}
		}
		f(1, 0)
		ans := a[0]
		if len(a[0]) > len(a[1]) {
			ans = a[1]
		}
		Fprintln(out, len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1176E(os.Stdin, os.Stdout) }
