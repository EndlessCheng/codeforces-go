package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 类似 https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/

// https://github.com/EndlessCheng
type node40 struct {
	lo, ro *node40
	cnt    int
}

func build40(l, r int) *node40 {
	o := &node40{}
	if l == r {
		return o
	}
	m := (l + r) >> 1
	o.lo = build40(l, m)
	o.ro = build40(m+1, r)
	return o
}

func (o node40) update(l, r, i int) *node40 {
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
	o.cnt = o.lo.cnt + o.ro.cnt
	return &o
}

func (o *node40) query(old *node40, l, r, k int) (int, int) {
	if l == r {
		return o.cnt - old.cnt, l
	}
	m := (l + r) >> 1
	cntL := o.lo.cnt - old.lo.cnt
	if k <= cntL {
		return o.lo.query(old.lo, l, m, k)
	}
	return o.ro.query(old.ro, m+1, r, k-cntL)
}

func cf840D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, v, l, r, k int
	Fscan(in, &n, &q)
	t := make([]*node40, n+1)
	t[0] = build40(1, n)
	for i := range n {
		Fscan(in, &v)
		t[i+1] = t[i].update(1, n, v)
	}
o:
	for range q {
		Fscan(in, &l, &r, &k)
		l--
		d := (r-l)/k + 1
		for k := d; k <= r-l; k += d {
			cnt, v := t[r].query(t[l], 1, n, k)
			if cnt >= d {
				Fprintln(out, v)
				continue o
			}
		}
		Fprintln(out, -1)
	}
}

//func main() { debug.SetGCPercent(-1); cf840D(bufio.NewReader(os.Stdin), os.Stdout) }
