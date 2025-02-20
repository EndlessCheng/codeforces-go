package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
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

func (t seg) update(o, i int, val int) {
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

func run(in io.Reader, out io.Writer) {
	const mx = 5e5
	f := make(seg, 2<<bits.Len(mx))
	f.build(1, 1, mx)
	var n, d, v int
	for Fscan(in, &n, &d); n > 0; n-- {
		Fscan(in, &v)
		f.update(1, v, f.query(1, max(v-d, 1), min(v+d, mx))+1)
	}
	Fprint(out, f[1].mx)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
