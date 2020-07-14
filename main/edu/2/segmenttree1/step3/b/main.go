package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type seg []struct{ l, r, sum int }

func (t seg) _pushUp(o int) { t[o].sum = t[o<<1].sum + t[o<<1|1].sum }

func (t seg) _build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = 1
		return
	}
	m := (l + r) >> 1
	t._build(o<<1, l, m)
	t._build(o<<1|1, m+1, r)
	t._pushUp(o)
}

func (t seg) _query(o, k int) (res int) {
	if t[o].l == t[o].r {
		t[o].sum = 0
		return t[o].l
	}
	defer t._pushUp(o)
	if k < t[o<<1|1].sum {
		return t._query(o<<1|1, k)
	}
	return t._query(o<<1, k-t[o<<1|1].sum)
}

func (t seg) init(n int)      { t._build(1, 1, n) }
func (t seg) query(k int) int { return t._query(1, k) }

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	t := make(seg, 4*n)
	t.init(n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for i := len(a) - 1; i >= 0; i-- {
		a[i] = t.query(a[i])
	}
	for _, v := range a {
		Fprint(out, v, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
