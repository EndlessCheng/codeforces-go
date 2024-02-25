package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1702E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v, w int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}
		for _, vs := range g {
			if len(vs) != 2 {
				Fprintln(out, "NO")
				continue o
			}
		}
		for v, vs := range g {
			if vs == nil {
				continue
			}
			cnt := 0
			pre := -1
			for g[v] != nil {
				cnt ^= 1
				if g[v][0] != pre {
					v, pre, g[v] = g[v][0], v, nil
				} else {
					v, pre, g[v] = g[v][1], v, nil
				}
			}
			if cnt > 0 {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { cf1702E(os.Stdin, os.Stdout) }
