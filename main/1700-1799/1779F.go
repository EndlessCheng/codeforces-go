package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf1779F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, p int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for w := 1; w < n; w++ {
		Fscan(in, &p)
		p--
		g[p] = append(g[p], w)
	}

	sz := make([]int, n)
	preF := make([]int, n)
	fs := make([]int, n)
	var dfs0 func(int)
	dfs0 = func(v int) {
		sz[v] = 1
		f := 1
		for _, w := range g[v] {
			preF[w] = f
			dfs0(w)
			sz[v] += sz[w]
			nf := 0
			for s := uint32(fs[w]); s > 0; s &= s - 1 {
				p := bits.TrailingZeros32(s)
				for t := uint32(f); t > 0; t &= t - 1 {
					q := bits.TrailingZeros32(t)
					nf |= 1 << (p ^ q)
				}
			}
			f = nf
		}

		nf := 0
		for s := uint32(f); s > 0; s &= s - 1 {
			p := bits.TrailingZeros32(s)
			nf |= 1 << (p ^ a[v])
		}
		f = nf

		if sz[v]%2 == 0 {
			f |= 1
		}

		fs[v] = f
	}
	dfs0(0)

	if fs[0]&1 == 0 {
		Fprint(out, -1)
		return
	}

	ans := []int{}
	var dfs func(int, int)
	dfs = func(v, tar int) {
		if tar == 0 && sz[v]%2 == 0 {
			ans = append(ans, v)
			return
		}
		tar ^= a[v]
		for i := len(g[v]) - 1; i >= 0; i-- {
			w := g[v][i]
			for s := uint32(fs[w]); s > 0; s &= s - 1 {
				p := bits.TrailingZeros32(s)
				if preF[w]>>(tar^p)&1 > 0 {
					tar ^= p
					dfs(w, p)
					break
				}
			}
		}
	}
	dfs(0, 0)

	ans = append(ans, 0)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v+1, " ")
	}
}

//func main() { cf1779F(bufio.NewReader(os.Stdin), os.Stdout) }
