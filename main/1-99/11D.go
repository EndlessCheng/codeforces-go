package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// github.com/EndlessCheng/codeforces-go
func CF11D(in io.Reader, out io.Writer) {
	var n, m, ans int
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

	f := make([][19]int, 1<<n)
	for i := range n {
		f[1<<i][i] = 1
	}
	for mask := range f {
		st := bits.TrailingZeros32(uint32(mask))
		for v, fv := range f[mask][:n] {
			if fv == 0 {
				continue
			}
			for _, w := range g[v] {
				if w == st {
					ans += fv
				} else if w > st && mask>>w&1 == 0 {
					f[mask|1<<w][w] += fv
				}
			}
		}
	}
	Fprint(out, (ans-m)/2)
}

//func main() { CF11D(os.Stdin, os.Stdout) }
