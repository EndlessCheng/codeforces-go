package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
const setInit1253, addInit1253 int = 1e18, 0

type seg1253 []struct{ l, r, max, set, add int }

func (t seg1253) doSet(O, v int) {
	o := &t[O]
	o.max = v
	o.set = v
	o.add = addInit1253
}

func (t seg1253) doAdd(O, v int) {
	o := &t[O]
	o.max += v
	if o.set != setInit1253 {
		o.set += v
	} else {
		o.add += v
	}
}

func (t seg1253) spread(o int) {
	if set := t[o].set; set != setInit1253 {
		t.doSet(o<<1, set)
		t.doSet(o<<1|1, set)
		t[o].set = setInit1253
	} else if add := t[o].add; add != addInit1253 {
		t.doAdd(o<<1, add)
		t.doAdd(o<<1|1, add)
		t[o].add = addInit1253
	}
}

func (t seg1253) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].set = setInit1253
	if l == r {
		t[o].max = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg1253) maintain(o int) {
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func (t seg1253) set(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.doSet(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.set(o<<1, l, r, v)
	}
	if m < r {
		t.set(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg1253) add(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.doAdd(o, v)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.add(o<<1, l, r, v)
	}
	if m < r {
		t.add(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg1253) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func p1253(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	ri := func() (x int) {
		neg := false
		b := rc()
		for ; '0' > b || b > '9'; b = rc() {
			if b == '-' {
				neg = true
			}
		}
		for ; '0' <= b && b <= '9'; b = rc() {
			x = x*10 + int(b&15)
		}
		if neg {
			return -x
		}
		return
	}

	n, q := ri(), ri()
	a := make([]int, n)
	for i := range a {
		a[i] = ri()
	}
	t := make(seg1253, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		op, l, r := ri(), ri(), ri()
		if op == 1 {
			t.set(1, l, r, ri())
		} else if op == 2 {
			t.add(1, l, r, ri())
		} else {
			Fprintln(out, t.query(1, l, r))
		}
	}
}

//func main() { p1253(os.Stdin, os.Stdout) }
