package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
const mod66 = 1_000_000_007
const mx66 = 6

var sPow66 [][mx66]int

type seg66 []struct {
	l, r int
	sum  [mx66]int
	todo int
}

func (seg66) mergeInfo(a, b [mx66]int) [mx66]int {
	for i, v := range b {
		a[i] += v // mod
	}
	return a
}

func (t seg66) apply(o, x int) {
	cur := &t[o]
	for i := range mx66 {
		cur.sum[i] = x * (sPow66[cur.r][i] - sPow66[cur.l-1][i]) % mod66
	}
	cur.todo = x
}

func (t seg66) maintain(o int) {
	t[o].sum = t.mergeInfo(t[o<<1].sum, t[o<<1|1].sum)
}

func (t seg66) spread(o int) {
	f := t[o].todo
	if f < 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = -1
}

func (t seg66) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = -1
	if l == r {
		t[o].sum[0] = a[l-1]
		for i := 1; i < mx66; i++ {
			t[o].sum[i] = t[o].sum[i-1] * l % mod66
		}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg66) update(o, l, r, x int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, x)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, x)
	}
	if m < r {
		t.update(o<<1|1, l, r, x)
	}
	t.maintain(o)
}

func (t seg66) query(o, l, r int) [mx66]int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf266E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	C := [mx66][mx66]int{}
	for i := range mx66 {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	var n, m, l, r, k int
	var op string
	Fscan(in, &n, &m)
	sPow66 = make([][mx66]int, n+1)
	for i := 1; i <= n; i++ {
		powI := 1
		for j := range mx66 {
			sPow66[i][j] = (sPow66[i-1][j] + powI) % mod66
			powI = powI * i % mod66
		}
	}

	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg66, 2<<bits.Len(uint(n)))
	t.build(a, 1, 1, n)

	for range m {
		Fscan(in, &op, &l, &r, &k)
		if op == "=" {
			t.update(1, l, r, k)
		} else {
			s := t.query(1, l, r)
			res := 0
			powL := 1
			for j := k; j >= 0; j-- {
				res += s[j] * C[k][j] % mod66 * powL
				powL = powL * -(l - 1) % mod66
			}
			Fprintln(out, (res%mod66+mod66)%mod66)
		}
	}
}

//func main() { cf266E(bufio.NewReader(os.Stdin), os.Stdout) }
