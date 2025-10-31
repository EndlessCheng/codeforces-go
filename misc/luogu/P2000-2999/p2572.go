package P2000_2999

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type info572 struct{ pre0, suf0, max0, cnt0, pre1, suf1, max1, cnt1 int }

type seg572 []struct {
	l, r int
	data info572
	set  int8
	flip int8
}

const setInit572, flipInit572 = -1, 0

func (seg572) mergeInfo(a, b info572) (c info572) {
	c.pre0 = a.pre0
	if a.cnt1 == 0 {
		c.pre0 += b.pre0
	}
	c.suf0 = b.suf0
	if b.cnt1 == 0 {
		c.suf0 += a.suf0
	}
	c.max0 = max(a.max0, b.max0, a.suf0+b.pre0)
	c.cnt0 = a.cnt0 + b.cnt0

	c.pre1 = a.pre1
	if a.cnt0 == 0 {
		c.pre1 += b.pre1
	}
	c.suf1 = b.suf1
	if b.cnt0 == 0 {
		c.suf1 += a.suf1
	}
	c.max1 = max(a.max1, b.max1, a.suf1+b.pre1)
	c.cnt1 = a.cnt1 + b.cnt1
	return
}

func (t seg572) doSet(O int, v int8) {
	o := &t[O]
	sz := o.r - o.l + 1
	if v == 0 {
		o.data = info572{sz, sz, sz, sz, 0, 0, 0, 0}
	} else {
		o.data = info572{0, 0, 0, 0, sz, sz, sz, sz}
	}
	o.set = v
	o.flip = flipInit572
}

func (t seg572) doFlip(O int) {
	o := &t[O]
	d := o.data
	o.data = info572{d.pre1, d.suf1, d.max1, d.cnt1, d.pre0, d.suf0, d.max0, d.cnt0}
	o.flip ^= 1
}

func (t seg572) spread(o int) {
	if v := t[o].set; v != setInit572 {
		t.doSet(o<<1, v)
		t.doSet(o<<1|1, v)
		t[o].set = setInit572
	}
	if t[o].flip != flipInit572 {
		t.doFlip(o << 1)
		t.doFlip(o<<1 | 1)
		t[o].flip = flipInit572
	}
}

func (t seg572) build(a []int8, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].set = setInit572
	if l == r {
		v := int(a[l-1])
		t[o].data = info572{v ^ 1, v ^ 1, v ^ 1, v ^ 1, v, v, v, v}
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg572) maintain(o int) {
	t[o].data = t.mergeInfo(t[o<<1].data, t[o<<1|1].data)
}

func (t seg572) update(o, l, r int, v int8) {
	if l <= t[o].l && t[o].r <= r {
		if v == 2 {
			t.doFlip(o)
		} else {
			t.doSet(o, v)
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t.maintain(o)
}

func (t seg572) query(o, l, r int) info572 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data
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

func p2572(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, q, l, r int
	var op int8
	Fscan(in, &n, &q)
	a := make([]int8, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg572, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for ; q > 0; q-- {
		Fscan(in, &op, &l, &r)
		l++
		r++
		switch op {
		case 3: Fprintln(out, t.query(1, l, r).cnt1)
		case 4: Fprintln(out, t.query(1, l, r).max1)
		default: t.update(1, l, r, op)
		}
	}
}

//func main() { p2572(os.Stdin, os.Stdout) }
