package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"sort"
)

// https://github.com/EndlessCheng
var lucky21 = []int{4, 7, 44, 47, 74, 77, 444, 447, 474, 477, 744, 747, 774, 777, 4444, 4447, 4474, 4477, 4744, 4747, 4774, 4777, 7444, 7447, 7474, 7477, 7744, 7747, 7774, 7777, 1e9}
var a21 []int

type seg21 []struct{ l, r, cnt, minD, todo int }

func (t seg21) set(o int) {
	i := t[o].l
	t[o].minD = lucky21[sort.SearchInts(lucky21, a21[i])] - a21[i]
	if t[o].minD == 0 {
		t[o].cnt = 1
	} else {
		t[o].cnt = 0
	}
}

func (t seg21) apply(o, f int) {
	t[o].minD -= f
	t[o].todo += f
}

func (t seg21) maintain(o int) {
	t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
	t[o].minD = min(t[o<<1].minD, t[o<<1|1].minD)
}

func (t seg21) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg21) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o)
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg21) update(o, l, r, f int) {
	if l <= t[o].l && t[o].r <= r && t[o].minD > f {
		t.apply(o, f)
		return
	}
	if t[o].l == t[o].r {
		a21[t[o].l] += t[o].todo + f
		t[o].todo = 0
		t.set(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
	t.maintain(o)
}

func (t seg21) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func cf121E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, d int
	var op string
	Fscan(in, &n, &m)
	a21 = make([]int, n)
	for i := range a21 {
		Fscan(in, &a21[i])
	}

	t := make(seg21, 2<<bits.Len(uint(n-1)))
	t.build(1, 0, n-1)

	for range m {
		Fscan(in, &op, &l, &r)
		l--
		r--
		if op[0] == 'a' {
			Fscan(in, &d)
			t.update(1, l, r, d)
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { cf121E(bufio.NewReader(os.Stdin), os.Stdout) }
