package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type data80 struct{ lc, rc int }
type seg80 []struct {
	l, r int
	data80
}

func (t seg80) do(l, r data80) data80 {
	min := l.lc
	if r.rc < min {
		min = r.rc
	}
	return data80{l.lc + r.lc - min, l.rc + r.rc - min}
}

func (t seg80) build(s string, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		if s[l-1] == '(' {
			t[o].lc = 1
		} else {
			t[o].rc = 1
		}
		return
	}
	m := (l + r) >> 1
	t.build(s, o<<1, l, m)
	t.build(s, o<<1|1, m+1, r)
	t[o].data80 = t.do(t[o<<1].data80, t[o<<1|1].data80)
}

func (t seg80) query(o, l, r int) data80 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].data80
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.do(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF380C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s string
	Fscan(in, &s)
	t := make(seg80, 4*len(s))
	t.build(s, 1, 1, len(s))
	var q, l, r int
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		d := t.query(1, l, r)
		Fprintln(out, r-l+1-d.lc-d.rc)
	}
}

//func main() { CF380C(os.Stdin, os.Stdout) }
