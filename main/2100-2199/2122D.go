package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf2122D(in io.Reader, out io.Writer) {
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		g := make([][]int, n)
		for range m {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		f := make([]int, n)
		for i := range f {
			f[i] = 1e9
		}
		f[0] = 0

		t := 0
		for ; f[n-1] >= 1e9; t++ {
			nf := slices.Clone(f)
			for i := range nf {
				nf[i]++
			}
			for v, ws := range g {
				w := ws[t%len(ws)]
				nf[w] = min(nf[w], f[v])
			}
			f = nf
		}
		Fprintln(out, t, f[n-1])
	}
}

//func main() { cf2122D(bufio.NewReader(os.Stdin), os.Stdout) }
