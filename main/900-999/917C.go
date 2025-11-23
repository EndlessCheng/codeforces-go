package main

import (
	. "fmt"
	"io"
	"math/bits"
	"slices"
)

// https://github.com/EndlessCheng
type matrix17 [][]int

func newMatrix17(n, m int) matrix17 {
	a := make(matrix17, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			a[i][j] = 1e18
		}
	}
	return a
}

func (a matrix17) mul(b matrix17) matrix17 {
	c := newMatrix17(len(a), len(b[0]))
	for i, row := range a {
		for k, x := range row {
			if x == 1e18 {
				continue
			}
			for j, y := range b[k] {
				c[i][j] = min(c[i][j], x+y)
			}
		}
	}
	return c
}

func (a matrix17) powMul(n int, f0 matrix17) matrix17 {
	res := f0
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = a.mul(res)
		}
		a = a.mul(a)
	}
	return res
}

func cf917C(in io.Reader, out io.Writer) {
	var x, k, n, q, t, ex int
	Fscan(in, &x, &k, &n, &q)
	c := make([]int, k)
	for i := range c {
		Fscan(in, &c[i])
	}

	// 不用 mp 是 1280ms，用 mp 是 280ms，刚好优化 1000ms！
	mp := make([]int, 1<<k)
	for i := range mp {
		if bits.OnesCount(uint(i)) == x {
			mp[i] = t
			t++
		}
	}

	m := newMatrix17(t, t)
	for old := range 1 << k {
		if bits.OnesCount(uint(old)) != x {
			continue
		}
		if old&1 == 0 {
			m[mp[old>>1]][mp[old]] = 0
			continue
		}
		nw := old >> 1
		for s := uint(1<<k - 1 ^ nw); s > 0; s &= s - 1 {
			p := bits.TrailingZeros(s)
			m[mp[nw|1<<p]][mp[old]] = c[p]
		}
	}

	f := newMatrix17(t, 1)
	f[0][0] = 0

	type pair struct{ p, w int }
	a := make([]pair, q)
	for i := range a {
		Fscan(in, &a[i].p, &a[i].w)
	}
	slices.SortFunc(a, func(a, b pair) int { return a.p - b.p })

	pre := 1
	for _, p := range a {
		if p.p > n-x {
			ex += p.w
			continue
		}
		f = m.powMul(p.p-pre, f)
		pre = p.p
		for i, idx := range mp {
			if i%2 > 0 && bits.OnesCount(uint(i)) == x {
				f[idx][0] += p.w
			}
		}
	}
	f = m.powMul(n-x+1-pre, f)
	Fprint(out, f[0][0]+ex)
}

//func main() { cf917C(bufio.NewReader(os.Stdin), os.Stdout) }
