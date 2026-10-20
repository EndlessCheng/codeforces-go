package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
var s string

type data struct{ rgb, rg, gb, r, g, b int }
type seg []data

func (seg) mergeData(a, b data) data {
	const mod = 998244353
	return data{
		(a.rgb + b.rgb + a.rg*b.b + a.r*b.gb) % mod,
		(a.rg + b.rg + a.r*b.g) % mod,
		(a.gb + b.gb + a.g*b.b) % mod,
		(a.r + b.r) % mod,
		(a.g + b.g) % mod,
		(a.b + b.b) % mod,
	}
}

func (t seg) maintain(node int) {
	t[node] = t.mergeData(t[node*2], t[node*2+1])
}

func (t seg) set(node, i, x int) {
	switch s[i] {
	case 'R':
		t[node].r = x
	case 'G':
		t[node].g = x
	default:
		t[node].b = x
	}
}

func (t seg) build(a []int, node, l, r int) {
	if l == r {
		t.set(node, l-1, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, node*2, l, m)
	t.build(a, node*2+1, m+1, r)
	t.maintain(node)
}

func (t seg) update(node, l, r, i, x int) {
	if l == r {
		t.set(node, l-1, x)
		return
	}
	m := (l + r) >> 1
	if i <= m {
		t.update(node*2, l, m, i, x)
	} else {
		t.update(node*2+1, m+1, r, i, x)
	}
	t.maintain(node)
}

func (t seg) query(node, l, r, ql, qr int) data {
	if ql <= l && r <= qr {
		return t[node]
	}
	m := (l + r) >> 1
	if qr <= m {
		return t.query(node*2, l, m, ql, qr)
	}
	if ql > m {
		return t.query(node*2+1, m+1, r, ql, qr)
	}
	return t.mergeData(t.query(node*2, l, m, ql, qr), t.query(node*2+1, m+1, r, ql, qr))
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, q, op, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fscan(in, &s)
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 1, n)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			t.update(1, 1, n, l, r)
		} else {
			Fprintln(out, t.query(1, 1, n, l, r).rgb)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
