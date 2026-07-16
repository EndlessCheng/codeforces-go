package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type seg []struct{ max, add int }

func (t seg) apply(o, f int) {
	t[o].max += f
	t[o].add += f
}

func (t seg) spread(o int) {
	f := t[o].add
	if f == 0 {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].add = 0
}

func (t seg) update(o, l, r, ql, qr, f int) {
	if ql <= l && r <= qr {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (l + r) >> 1
	if ql <= m {
		t.update(o<<1, l, m, ql, qr, f)
	}
	if m < qr {
		t.update(o<<1|1, m+1, r, ql, qr, f)
	}
	t[o].max = max(t[o<<1].max, t[o<<1|1].max)
}

func run(in io.Reader, out io.Writer) {
	var n, v int
	Fscan(in, &n)
	f := [3]seg{}
	for k := range f {
		f[k] = make(seg, 2<<bits.Len(uint(n-1)))
	}
	last := make([]int, n+1)

	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		for k := 2; k >= 0; k-- {
			if k > 0 {
				f[k].update(1, 1, n, i, i, f[k-1][1].max)
			}
			f[k].update(1, 1, n, last[v]+1, i, 1)
		}
		last[v] = i
	}
	Fprint(out, f[2][1].max)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
