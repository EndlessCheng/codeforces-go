package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
type seg05 []struct {
	l, r  int
	state int8
	flip  bool
}

func (t seg05) maintain(o int) {
	x, y := t[o<<1].state, t[o<<1|1].state
	if x < 0 && y < 0 {
		t[o].state = -1
	} else if x > 0 && y > 0 {
		t[o].state = 1
	} else {
		t[o].state = 0
	}
}

func (t seg05) build(o, l, r int) {
	t[o].l, t[o].r, t[o].state = l, r, -1
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg05) do(o int) {
	t[o].state = -t[o].state
	t[o].flip = !t[o].flip
}

func (t seg05) spread(o int) {
	if t[o].flip {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
}

func (t seg05) flip(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.flip(o<<1, l, r)
	}
	if m < r {
		t.flip(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg05) next0(o, l int) int {
	if t[o].l == t[o].r {
		if t[o].state < 0 {
			return t[o].l
		}
		return 0
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m && t[o<<1].state <= 0 {
		if p := t.next0(o<<1, l); p > 0 {
			return p
		}
	}
	return t.next0(o<<1|1, l)
}

func (t seg05) next1(o, l int) int {
	if t[o].l == t[o].r {
		if t[o].state > 0 {
			return t[o].l
		}
		return 0
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m && t[o<<1].state >= 0 {
		if p := t.next1(o<<1, l); p > 0 {
			return p
		}
	}
	return t.next1(o<<1|1, l)
}

func (t seg05) last1(o int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	t.spread(o)
	if t[o<<1|1].state >= 0 {
		return t.last1(o<<1 | 1)
	}
	return t.last1(o << 1)
}

func CF1705E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mx int = 2e5 + 20

	var n, q, i, v int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	t := make(seg05, mx*4)
	t.build(1, 1, mx)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		t.flip(1, a[i], t.next0(1, a[i]))
	}
	for ; q > 0; q-- {
		Fscan(in, &i, &v)
		t.flip(1, a[i], t.next1(1, a[i]))
		a[i] = v
		t.flip(1, v, t.next0(1, v))
		Fprintln(out, t.last1(1))
	}
}

//func main() { CF1705E(os.Stdin, os.Stdout) }
