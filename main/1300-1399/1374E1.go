package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1374E1(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	var n, k, t, p, q, s int
	g := [4][]int{}
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &t, &p, &q)
		x := p<<1 | q
		g[x] = append(g[x], t)
	}
	a, b, both := g[1], g[2], g[3]
	if len(a) > len(b) {
		a, b = b, a
	}
	na, nb := len(a), len(both)
	if na+nb < k {
		Fprint(out, -1)
		return
	}

	sort.Ints(a)
	sort.Ints(b)
	if na > k {
		a = a[:k]
		na = k
	}
	for i, v := range b[:na] {
		a[i] += v
	}

	sort.Ints(both)
	if nb > k {
		both = both[:k]
		nb = k
	}

	for _, v := range both {
		s += v
	}
	for _, v := range a[:k-nb] {
		s += v
	}
	ans := s

	for i, v := range a[k-nb:] {
		s += v - both[nb-1-i]
		ans = min(ans, s)
	}
	Fprint(out, ans)
}

//func main() { CF1374E1(os.Stdin, os.Stdout) }
