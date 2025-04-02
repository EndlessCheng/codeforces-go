package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, m, k, w, v, c int
	Fscan(in, &n, &m, &k)
	type pair struct{ w, v int }
	g := make([][]pair, n)
	for i := 0; i < n; i++ {
		Fscan(in, &w, &v, &c)
		g[c-1] = append(g[c-1], pair{w, v})
	}

	f := make([]int, m+1)
	for _, a := range g {
		nf := append(f[:0:0], f...)
		for _, p := range a {
			for j := m; j >= p.w; j-- {
				nf[j] = max(nf[j], f[j-p.w]+p.v+k, nf[j-p.w]+p.v)
			}
		}
		f = nf
	}
	Fprint(out, f[m])
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
