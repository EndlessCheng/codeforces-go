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

	var n, k, t, b0, b1, s int
	g := [4][]int{}
	for Fscan(in, &n, &k); n > 0; n-- {
		Fscan(in, &t, &b0, &b1)
		m := b0<<1 | b1
		g[m] = append(g[m], t)
	}
	a, b, both := g[1], g[2], g[3]
	if len(a) > len(b) {
		a, b = b, a
	}
	n, m := len(a), len(both)
	if n+m < k {
		Fprint(out, -1)
		return
	}

	sort.Ints(a)
	sort.Ints(b)
	for i, v := range b[:n] {
		a[i] += v
	}
	sort.Ints(both)
	if m > k {
		both = both[:k]
		m = k
	}

	for _, v := range both {
		s += v
	}
	for _, v := range a[:k-m] {
		s += v
	}
	ans := s
	for i := k - m; i < n; i++ {
		s += a[i]
		if k-1-i >= 0 {
			s -= both[k-1-i]
		}
		ans = min(ans, s)
	}
	Fprint(out, ans)
}

//func main() { CF1374E1(os.Stdin, os.Stdout) }
