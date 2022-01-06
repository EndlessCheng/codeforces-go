package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type seg55 []struct{ l, r, min, todo int }

func (t seg55) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg55) do(o, v int) {
	t[o].min += v
	t[o].todo += v
}

func (t seg55) update(o, l, r, v int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, v)
		return
	}
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, v)
	}
	if m < r {
		t.update(o<<1|1, l, r, v)
	}
	t[o].min = min55(t[o<<1].min, t[o<<1|1].min)
}

func CF1555E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, l int
	Fscan(in, &n, &m)
	a := make([]struct{ l, r, w int }, n)
	for i := range a {
		Fscan(in, &a[i].l, &a[i].r, &a[i].w)
		a[i].r--
	}
	sort.Slice(a, func(i, j int) bool { return a[i].w < a[j].w })

	t := make(seg55, m*4)
	t.build(1, 1, m-1)
	ans := int(1e9)
	for _, p := range a {
		t.update(1, p.l, p.r, 1)
		for t[1].min > 0 {
			ans = min55(ans, p.w-a[l].w)
			t.update(1, a[l].l, a[l].r, -1)
			l++
		}
	}
	Fprint(out, ans)
}

//func main() { CF1555E(os.Stdin, os.Stdout) }

func min55(a, b int) int {
	if a > b {
		return b
	}
	return a
}
