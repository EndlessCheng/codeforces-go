package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
type data39 struct{ lv, rv, preD, preID, sufI, sufID, mx int }
type seg39 []struct {
	l, r int
	d    data39
	todo int
}

func (t seg39) apply(o, f int) {
	cur := &t[o]
	cur.d.lv += f
	cur.d.rv += f
	cur.todo += f
}

func (t seg39) maintain(o int) {
	lo, ro := &t[o<<1], &t[o<<1|1]
	lsz := lo.r - lo.l + 1
	rsz := ro.r - ro.l + 1
	l, r := lo.d, ro.d
	inc := l.rv < r.lv
	dec := l.rv > r.lv
	ne := l.rv != r.lv

	preD := l.preD
	if dec && preD == lsz {
		preD += r.preD
	}
	sufI := r.sufI
	if inc && sufI == rsz {
		sufI += l.sufI
	}

	preID := l.preID
	if dec && preID == lsz {
		preID += r.preD
	}
	if l.sufI == lsz {
		if ne {
			preID = max(preID, lsz+r.preD)
		}
		if inc {
			preID = max(preID, lsz+r.preID)
		}
	}

	sufID := r.sufID
	if inc && sufID == rsz {
		sufID += l.sufI
	}
	if r.preD == rsz {
		if ne {
			sufID = max(sufID, l.sufI+rsz)
		}
		if dec {
			sufID = max(sufID, l.sufID+rsz)
		}
	}

	mx := max(l.mx, r.mx)
	if ne {
		mx = max(mx, l.sufI+r.preD)
	}
	if dec {
		mx = max(mx, l.sufID+r.preD)
	}
	if inc {
		mx = max(mx, l.sufI+r.preID)
	}

	t[o].d = data39{l.lv, r.rv, preD, preID, sufI, sufID, mx}
}

func (t seg39) spread(o int) {
	f := t[o].todo
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = 0
}

func (t seg39) build(in io.Reader, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		var v int
		Fscan(in, &v)
		t[o].d = data39{v, v, 1, 1, 1, 1, 1}
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg39) update(o, l, r, f int) {
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

func cf739C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, v int
	Fscan(in, &n)
	t := make(seg39, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)
	Fscan(in, &m)
	for range m {
		Fscan(in, &l, &r, &v)
		t.update(1, l, r, v)
		Fprintln(out, t[1].d.mx)
	}
}

//func main() { cf739C(bufio.NewReader(os.Stdin), os.Stdout) }
