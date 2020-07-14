package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct {
	l, r int
	inv  int64
	cnt  [41]int
}

func (t seg) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	inv := lo.inv + ro.inv
	s := 0
	for i := 40; i > 1; i-- {
		s += lo.cnt[i]
		lo.cnt[i] += ro.cnt[i]
		inv += int64(s) * int64(ro.cnt[i-1])
	}
	lo.cnt[1] += ro.cnt[1]
	t[o].inv = inv
	t[o].cnt = lo.cnt
}

func (t seg) _build(a []int8, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].cnt[a[l]]++
		return
	}
	m := (l + r) >> 1
	t._build(a, o<<1, l, m)
	t._build(a, o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _update(o, i int, pre, cur int8) {
	if t[o].l == t[o].r {
		t[o].cnt[pre]--
		t[o].cnt[cur]++
		return
	}
	if i <= (t[o].l+t[o].r)>>1 {
		t._update(o<<1, i, pre, cur)
	} else {
		t._update(o<<1|1, i, pre, cur)
	}
	t._pushUp(o)
}

func (t seg) _query(o, l, r int) (int64, [41]int) {
	if l <= t[o].l && t[o].r <= r {
		return t[o].inv, t[o].cnt
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t._query(o<<1, l, r)
	}
	if l > m {
		return t._query(o<<1|1, l, r)
	}
	invL, cntL := t._query(o<<1, l, r)
	invR, cntR := t._query(o<<1|1, l, r)
	inv := invL + invR
	s := 0
	for i := 40; i > 1; i-- {
		s += cntL[i]
		cntL[i] += cntR[i]
		inv += int64(s) * int64(cntR[i-1])
	}
	cntL[1] += cntR[1]
	return inv, cntL
}

func (t seg) init(a []int8)               { t._build(a, 1, 1, len(a)-1) }
func (t seg) update(i int, pre, cur int8) { t._update(1, i, pre, cur) }
func (t seg) query(l, r int) int64        { inv, _ := t._query(1, l, r); return inv }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, op, i, l, r int
	var v int8
	Fscan(in, &n, &q)
	a := make([]int8, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	t := make(seg, 4*len(a))
	t.init(a)
	for ; q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			Fscan(in, &l, &r)
			Fprintln(out, t.query(l, r))
		} else {
			Fscan(in, &i, &v)
			t.update(i, a[i], v)
			a[i] = v
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
