package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1200F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 2520

	var n, m, t, v, c int
	Fscan(in, &n)
	k := make([]int, n)
	for i := range k {
		Fscan(in, &k[i])
		k[i] = (k[i]%mod + mod) % mod
	}
	to := make([][]int, n)
	for i := range to {
		Fscan(in, &m)
		to[i] = make([]int, m)
		for j := range to[i] {
			Fscan(in, &to[i][j])
			to[i][j]--
		}
	}

	n *= 2520
	g := make([]int, n)
	deg := make([]int, n)
	for i, v := range k {
		for j := 0; j < 2520; j++ {
			c := (j + v) % 2520
			w := to[i][c%len(to[i])]*2520 + c
			g[i*2520+j] = w
			deg[w]++
		}
	}

	q := []int{}
	for i, d := range deg {
		if d == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w := g[v]
		if deg[w]--; deg[w] == 0 {
			q = append(q, w)
		}
	}

	dpL := make([]int, n)
	var f func(int) int
	f = func(v int) int {
		if dpL[v] > 0 {
			return dpL[v]
		}
		if deg[v] == 0 {
			dpL[v] = f(g[v])
		} else {
			mp := map[int]bool{}
			for w := v; ; w = g[w] {
				mp[v/2520] = true
				if g[w] == v {
					break
				}
			}
			for w := v; ; w = g[w] {
				dpL[w] = len(mp)
				if g[w] == v {
					break
				}
			}
		}
		return dpL[v]
	}

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &v, &c)
		Fprintln(out, f((c%mod+mod)%mod+(v-1)*2520))
	}
}

//func main() { CF1200F(os.Stdin, os.Stdout) }
