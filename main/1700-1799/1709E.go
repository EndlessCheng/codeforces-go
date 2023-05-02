package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1709E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v, w int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	ans := 0
	var f func(int, int, int) map[int]bool
	f = func(v, fa, xor int) map[int]bool {
		xor ^= a[v]
		m := map[int]bool{xor: true}
	o:
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			m2 := f(w, v, xor)
			if m == nil {
				continue
			}
			if len(m2) > len(m) {
				m, m2 = m2, m
			}
			for x := range m2 {
				if m[x^a[v]] {
					ans++
					m = nil
					continue o
				}
			}
			for x := range m2 {
				m[x] = true
			}
		}
		return m
	}
	f(0, -1, 0)
	Fprint(out, ans)
}

//func main() { CF1709E(os.Stdin, os.Stdout) }
