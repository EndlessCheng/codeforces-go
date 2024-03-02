package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://space.bilibili.com/206214
type seg36 []struct{ l, r, min int }

func (t seg36) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg36) update(o, i, val int) {
	if t[o].l == t[o].r {
		t[o].min = val
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t[o].min = min(t[o<<1].min, t[o<<1|1].min)
}

func (t seg36) preMin(o, r int) int {
	if t[o].r <= r {
		return t[o].min
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.preMin(o<<1, r)
	}
	return min(t[o<<1].min, t.preMin(o<<1|1, r))
}

func cf1436E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, v int
	Fscan(in, &n)
	has := make([]bool, n+3)
	last := make([]int, n+2)
	t := make(seg36, 2<<bits.Len(uint(n-1)))
	t.build(1, 1, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if v == 1 {
			has[2] = true
		} else {
			has[1] = true
			if t.preMin(1, v-1) > last[v] {
				has[v] = true
			}
			last[v] = i
		}
		t.update(1, v, i)
	}
	for v := 2; v <= n+1; v++ {
		if t.preMin(1, v-1) > last[v] {
			has[v] = true
		}
	}
	mex := 1
	for has[mex] {
		mex++
	}
	Fprint(out, mex)
}

//func main() { cf1436E(os.Stdin, os.Stdout) }
