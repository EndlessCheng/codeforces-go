package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
type seg33 []struct {
	l, r      int
	max, todo int
}

func (t seg33) maintain(o int) {
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func (t seg33) do(o, v int) {
	t[o].max += v
	t[o].todo += v
}

func (t seg33) build(f []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = 0
	if l == r {
		t[o].max = f[l-1] // 注意这里是 f[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(f, o<<1, l, m)
	t.build(f, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg33) spread(o int) {
	if v := t[o].todo; v != 0 {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = 0
	}
}

func (t seg33) inc(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, 1)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.inc(o<<1, l, r)
	}
	if m < r {
		t.inc(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg33) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].max
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return max(t.query(o<<1, l, r), t.query(o<<1|1, l, r))
}

func CF833B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k, v int
	Fscan(in, &n, &k)
	pre := make([]int, n+1)
	p := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		pre[i] = p[v]
		p[v] = i
	}

	f := make([]int, n+1)
	t := make(seg33, n*4)
	for ; k > 0; k-- {
		t.build(f, 1, 1, n)
		for i := 1; i <= n; i++ {
			t.inc(1, pre[i]+1, i)
			f[i] = t.query(1, 1, i)
		}
	}
	Fprint(out, f[n])
}

//func main() { CF833B(os.Stdin, os.Stdout) }
