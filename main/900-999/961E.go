package main

import (
	. "fmt"
	"io"
	"runtime/debug"
)

// https://github.com/EndlessCheng
func init() { debug.SetGCPercent(-1) }

type data61 struct{ cnt int }
type node61 struct {
	lo, ro *node61
	data61
}

func (node61) merge(l, r data61) data61 {
	return data61{l.cnt + r.cnt}
}

func (o *node61) maintain() {
	o.data61 = o.merge(o.lo.data61, o.ro.data61)
}

func build61(l, r int) *node61 {
	o := &node61{}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build61(l, m)
	o.ro = build61(m+1, r)
	return o
}

func (o node61) update(l, r, i int) *node61 {
	if l == r {
		o.cnt++
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		o.lo = o.lo.update(l, m, i)
	} else {
		o.ro = o.ro.update(m+1, r, i)
	}
	o.maintain()
	return &o
}

func (o *node61) query(l, r, ql, qr int) data61 {
	if ql <= l && r <= qr {
		return o.data61
	}
	m := (l + r) >> 1
	if qr <= m {
		return o.lo.query(l, m, ql, qr)
	}
	if m < ql {
		return o.ro.query(m+1, r, ql, qr)
	}
	return o.merge(o.lo.query(l, m, ql, qr), o.ro.query(m+1, r, ql, qr))
}

func cf961E(in io.Reader, out io.Writer) {
	var n, v, ans int
	Fscan(in, &n)
	t := make([]*node61, n+1)
	t[0] = build61(0, n)
	for i := range n {
		Fscan(in, &v)
		r := min(v, i)
		ans += r - t[r].query(0, n, 0, i).cnt
		t[i+1] = t[i].update(0, n, min(v, n))
	}
	Fprint(out, ans)
}

//func main() { cf961E(bufio.NewReader(os.Stdin), os.Stdout) }
