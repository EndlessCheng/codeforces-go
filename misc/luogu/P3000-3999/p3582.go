package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg582 []struct{ max, add int }

func (t seg582) apply(o, x int) {
	t[o].max += x
	t[o].add += x
}

func (t seg582) spread(o int) {
	f := t[o].add
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].add = 0
}

func (t seg582) maintain(o int) {
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func (t seg582) update(o, l, r, ql, qr, x int) {
	if ql <= l && r <= qr {
		t.apply(o, x)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, x)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, x)
	}
	t.maintain(o)
}

func p3582(in io.Reader, out io.Writer) {
	var n, m int
	Fscan(in, &n, &m)
	p := make([]int, n)
	for i := range p {
		Fscan(in, &p[i])
	}
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}

	t := make(seg582, 2<<bits.Len(uint(n-1)))
	last := make([]int, m)
	last2 := make([]int, m)
	ans := 0
	for i, v := range p {
		v--
		l := last[v]
		t.update(1, 1, n, l+1, i+1, a[v])
		if l > 0 {
			t.update(1, 1, n, last2[v]+1, l, -a[v])
			last2[v] = l
		}
		last[v] = i + 1
		ans = max(ans, t[1].max)
	}
	Fprint(out, ans)
}

//func main() { p3582(bufio.NewReader(os.Stdin), os.Stdout) }
