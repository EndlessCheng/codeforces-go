package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type data01 struct{ cnt, sumC, sumC2 int }
type seg01 []struct {
	l, r int
	data01
	todo int
}

func (seg01) merge(l, r data01) data01 {
	return data01{l.cnt + r.cnt, l.sumC + r.sumC, l.sumC2 + r.sumC2}
}

func (t seg01) apply(o, f int) {
	cur := &t[o]
	cur.sumC2 += cur.sumC*f*2 + cur.cnt*f*f
	cur.sumC += cur.cnt * f
	cur.todo += f
}

func (t seg01) maintain(o int) {
	t[o].data01 = t.merge(t[o<<1].data01, t[o<<1|1].data01)
}

func (t seg01) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg01) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg01) set(o, i, cnt int) {
	if t[o].l == t[o].r {
		if cnt < 0 {
			t[o].data01 = data01{}
		} else {
			t[o].data01 = data01{1, cnt, cnt * cnt}
		}
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.set(o<<1, i, cnt)
	} else {
		t.set(o<<1|1, i, cnt)
	}
	t.maintain(o)
}

func (t seg01) update(o, l, r, f int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
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

func (t seg01) query(o, l, r int) data01 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data01
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.merge(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func cf1701F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	const mx int = 2e5
	t := make(seg01, 2<<bits.Len(uint(mx)))
	t.build(1, 1, mx)
	has := [mx + 1]bool{}

	var q, d, i int
	Fscan(in, &q, &d)
	for range q {
		Fscan(in, &i)
		if !has[i] {
			has[i] = true
			cnt := t.query(1, i, min(i+d, mx)).cnt
			t.set(1, i, cnt)
			if i > 1 {
				t.update(1, max(i-d, 1), i-1, 1)
			}
		} else {
			has[i] = false
			t.set(1, i, -1)
			if i > 1 {
				t.update(1, max(i-d, 1), i-1, -1)
			}
		}
		Fprintln(out, (t[1].sumC2-t[1].sumC)/2)
	}
}

//func main() { cf1701F(bufio.NewReader(os.Stdin), os.Stdout) }
