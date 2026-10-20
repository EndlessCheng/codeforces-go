package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
type data struct {
	inv int
	cnt [3]int
}

type seg []data

func (seg) mergeData(a, b data) data {
	return data{
		a.inv + b.inv + a.cnt[1]*b.cnt[0] + a.cnt[2]*(b.cnt[0]+b.cnt[1]),
		[3]int{a.cnt[0] + b.cnt[0], a.cnt[1] + b.cnt[1], a.cnt[2] + b.cnt[2]},
	}
}

func (t seg) maintain(node int) {
	t[node] = t.mergeData(t[node*2], t[node*2+1])
}

func (t seg) build(in io.Reader, node, l, r int) {
	if l == r {
		var x int
		Fscan(in, &x)
		t[node].cnt[x] = 1
		return
	}
	m := (l + r) >> 1
	t.build(in, node*2, l, m)
	t.build(in, node*2+1, m+1, r)
	t.maintain(node)
}

func (t seg) update(node, l, r, i, x int) {
	if l == r {
		t[node].cnt = [3]int{}
		t[node].cnt[x] = 1
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
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 1, n)
	for range q {
		Fscan(in, &op, &l, &r)
		if op == 1 {
			t.update(1, 1, n, l, r)
		} else {
			Fprintln(out, t.query(1, 1, n, l, r).inv)
		}
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
