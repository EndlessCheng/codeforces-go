package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
	"slices"
)

// https://github.com/EndlessCheng
type seg []int

func (t seg) update(o, l, r, i, v int) {
	if l == r {
		t[o] = v
		return
	}
	m := (l + r) / 2
	if i <= m {
		t.update(o*2, l, m, i, v)
	} else {
		t.update(o*2+1, m+1, r, i, v)
	}
	t[o] = max(t[o*2], t[o*2+1])
}

func (t seg) query(o, l, r, ql, qr int) int {
	if ql <= l && r <= qr {
		return t[o]
	}
	m := (l + r) / 2
	if qr <= m {
		return t.query(o*2, l, m, ql, qr)
	}
	if ql > m {
		return t.query(o*2+1, m+1, r, ql, qr)
	}
	return max(t.query(o*2, l, m, ql, qr), t.query(o*2+1, m+1, r, ql, qr))
}

func run(in io.Reader, out io.Writer) {
	var n, d, r, v int
	Fscan(in, &n, &d, &r)
	pos := make([]int, n+1)
	for i := range n {
		Fscan(in, &v)
		pos[v] = i
	}

	t := make(seg, 2<<bits.Len(uint(n-1)))
	for i := range t {
		t[i] = -1e9
	}
	f := make([]int, n)
	for x := d + 1; x <= n; x++ {
		t.update(1, 0, n-1, pos[x-d], f[pos[x-d]])
		i := pos[x]
		f[i] = max(t.query(1, 0, n-1, max(i-r, 0), min(i+r, n-1))+1, 0)
	}
	Fprint(out, slices.Max(f))
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
