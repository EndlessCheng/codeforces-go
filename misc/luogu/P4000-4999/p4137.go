package main

import (
	"bufio"
	. "fmt"
	"io"
	"runtime/debug"
)

// https://space.bilibili.com/206214
func init() { debug.SetGCPercent(-1) }

type node137 struct {
	lo, ro *node137
	mn     int
}

func build137(l, r int) *node137 {
	o := &node137{mn: -1}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build137(l, m)
	o.ro = build137(m+1, r)
	return o
}

func (o node137) update(l, r, i, val int) *node137 {
	if l == r {
		o.mn = val
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		o.lo = o.lo.update(l, m, i, val)
	} else {
		o.ro = o.ro.update(m+1, r, i, val)
	}
	o.mn = min(o.lo.mn, o.ro.mn)
	return &o
}

func (o *node137) query(l, r, ql int) int {
	if l == r {
		if o.mn >= ql { // 元素 l 在询问区间内
			return l + 1
		}
		return l
	}
	m := (l + r) >> 1
	if o.lo.mn < ql {
		return o.lo.query(l, m, ql)
	}
	return o.ro.query(m+1, r, ql)
}

func p4137(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, v, l, r int
	Fscan(in, &n, &m)

	t := make([]*node137, n+1)
	t[0] = build137(0, n-1)
	for i := range n {
		Fscan(in, &v)
		if v >= n {
			t[i+1] = t[i]
		} else {
			t[i+1] = t[i].update(0, n-1, v, i)
		}
	}

	for range m {
		Fscan(in, &l, &r)
		Fprintln(out, t[r].query(0, n-1, l-1))
	}
}

//func main() { p4137(bufio.NewReader(os.Stdin), os.Stdout) }
