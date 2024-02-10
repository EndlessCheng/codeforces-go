package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"os"
	"runtime/debug"
	"time"
)

// https://space.bilibili.com/206214
var null = &node{}

func init() {
	debug.SetGCPercent(-1)
	null.lo = null
	null.ro = null
}

type node struct {
	lo, ro *node
	sum    int
}

func (o node) modify(l, r, i int) *node {
	if l == r {
		o.sum += i
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		if o.lo == null {
			o.lo = &node{lo: null, ro: null}
		}
		o.lo = o.lo.modify(l, m, i)
	} else {
		if o.ro == null {
			o.ro = &node{lo: null, ro: null}
		}
		o.ro = o.ro.modify(m+1, r, i)
	}
	o.sum = o.lo.sum + o.ro.sum
	return &o
}

func (o *node) query(old *node, l, r, upper int) int {
	if l == r {
		return o.sum - old.sum
	}
	m := (l + r) >> 1
	if upper <= m {
		return o.lo.query(old.lo, l, m, upper)
	}
	return o.lo.sum - old.lo.sum + o.ro.query(old.ro, m+1, r, upper)
}

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	rand.Seed(time.Now().UnixNano())
	const R int = 1e9

	var n, v, q, l, r, ans int
	Fscan(in, &n)
	t := make([]*node, n+1)
	t[0] = &node{lo: null, ro: null}
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		t[i] = t[i-1].modify(0, R, v)
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r, &v)
		l ^= ans
		r ^= ans
		v ^= ans
		ans = t[r].query(t[l-1], 0, R, v)
		Fprintln(out, ans)
	}
}

func main() { run(os.Stdin, os.Stdout) }
