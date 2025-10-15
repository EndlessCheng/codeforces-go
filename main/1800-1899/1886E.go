package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
func cf1886E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, b int
	Fscan(in, &n, &m)
	type pair struct{ v, i int }
	a := make([]pair, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i
	}
	slices.SortFunc(a, func(a, b pair) int { return b.v - a.v })

	minR := make([][]int, m)
	for i := range minR {
		Fscan(in, &b)
		mn := make([]int, n)
		for j, p := range a {
			mn[j] = 1e9
			need := (b-1)/p.v + 1
			if l := j - need + 1; l >= 0 {
				mn[l] = min(mn[l], j)
			}
		}
		for j := n - 2; j >= 0; j-- {
			mn[j] = min(mn[j], mn[j+1])
		}
		minR[i] = mn
	}

	u := 1 << m
	f := make([]int, u)
	for i := range f {
		f[i] = 1e9
	}
	f[0] = 0
	from := make([]int, u)
	for s, man := range f {
		if man >= n {
			continue
		}
		for t, lb := u-1^s, 0; t > 0; t ^= lb {
			lb = t & -t
			ns := s | lb
			i := bits.TrailingZeros32(uint32(lb))
			need := minR[i][man] - man + 1
			if man+need < f[ns] {
				f[ns] = man + need
				from[ns] = s
			}
		}
	}
	if f[u-1] > n {
		Fprint(out, "NO")
		return
	}

	type pair2 struct{ l, r int }
	ans := make([]pair2, m)
	for i := u - 1; i > 0; i = from[i] {
		j := from[i]
		ans[bits.TrailingZeros32(uint32(i^j))] = pair2{f[j], f[i]}
	}

	Fprintln(out, "YES")
	for _, p := range ans {
		Fprint(out, p.r-p.l)
		for _, q := range a[p.l:p.r] {
			Fprint(out, " ", q.i+1)
		}
		Fprintln(out)
	}
}

//func main() { cf1886E(bufio.NewReader(os.Stdin), os.Stdout) }
