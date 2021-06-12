package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type seg []struct{ l, r, mx int }

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].mx = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t[o].mx = max(t[o<<1].mx, t[o<<1|1].mx)
}

func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mx
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans int
	Fscan(in, &n)
	a := make([]struct{ v, i int }, n)
	for i := range a {
		Fscan(in, &a[i].v)
		a[i].i = i + 1
	}
	sort.Slice(a, func(i, j int) bool { a, b := a[i], a[j]; return a.v < b.v || a.v == b.v && a.i > b.i })
	t := make(seg, 4*n)
	t.build(1, 1, n)
	for _, p := range a {
		v := p.v + t.query(1, 1, p.i)
		ans = max(ans, v)
		t.update(1, p.i, v)
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
