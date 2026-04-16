package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type pair struct{ sum, gcd int }

type seg []struct {
	l, r int
	pair
}

func mergeInfo(l, r pair) pair {
	return pair{l.sum + r.sum, gcd463(l.gcd, r.gcd)}
}

func (t seg) maintain(o int) {
	t[o].pair = mergeInfo(t[o<<1].pair, t[o<<1|1].pair)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].pair = pair{a[l-1], a[l-1]}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) update(o, i, v int) {
	cur := &t[o]
	if cur.l == cur.r {
		cur.sum += v
		cur.gcd += v
		return
	}
	m := (cur.l + cur.r) >> 1
	if i <= m {
		t.update(o<<1, i, v)
	} else {
		t.update(o<<1|1, i, v)
	}
	t.maintain(o)
}

func (t seg) query(o, l, r int) pair {
	if l <= t[o].l && t[o].r <= r {
		return t[o].pair
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return mergeInfo(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p10463(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	
	var n, m, l, r, d int
	var op string
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := n - 1; i > 0; i-- {
		a[i] -= a[i-1]
	}

	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)

	for range m {
		Fscan(in, &op, &l, &r)
		if op == "C" {
			Fscan(in, &d)
			t.update(1, l, d)
			if r < n {
				t.update(1, r+1, -d)
			}
		} else {
			ans := t.query(1, 1, l).sum
			if l < r {
				g := t.query(1, l+1, r).gcd
				ans = abs(gcd463(ans, g))
			}
			Fprintln(out, ans)
		}
	}
}

//func main() { p10463(bufio.NewReader(os.Stdin), os.Stdout) }
func gcd463(a, b int) int { for a != 0 { a, b = b%a, a }; return b }
