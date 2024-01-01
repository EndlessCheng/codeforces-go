package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1790F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		c := make([]int, n)
		for i := range c {
			Fscan(in, &c[i])
			c[i]--
		}
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		pa := make([]int, n)
		pa[0] = -1
		var f func(int)
		f = func(v int) {
			for _, w := range g[v] {
				if w != pa[v] {
					pa[w] = v
					f(w)
				}
			}
		}
		f(0)

		mn := n + 1
		dis := make([]int, n)
		for i := range dis {
			dis[i] = n + 1
		}
		for i, v := range c {
			for step := 0; v >= 0 && step < mn; step++ {
				mn = min(mn, dis[v]+step)
				dis[v] = min(dis[v], step)
				v = pa[v]
			}
			if i > 0 {
				Fprint(out, mn, " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { CF1790F(os.Stdin, os.Stdout) }
