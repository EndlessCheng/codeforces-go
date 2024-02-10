package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand"
	"runtime/debug"
	"time"
)

// https://space.bilibili.com/206214
var null = &node71{}

func init() {
	debug.SetGCPercent(-1)
	null.lo = null
	null.ro = null
}

type node71 struct {
	lo, ro *node71
	xorS   int
}

func (o node71) modify(l, r, i int, v int) *node71 {
	if l == r {
		o.xorS ^= v
		return &o
	}
	m := (l + r) >> 1
	if i <= m {
		if o.lo == null {
			o.lo = &node71{lo: null, ro: null}
		}
		o.lo = o.lo.modify(l, m, i, v)
	} else {
		if o.ro == null {
			o.ro = &node71{lo: null, ro: null}
		}
		o.ro = o.ro.modify(m+1, r, i, v)
	}
	o.xorS = o.lo.xorS ^ o.ro.xorS
	return &o
}

func (o *node71) query(old *node71, l, r int) int {
	if l == r {
		return l
	}
	m := (l + r) >> 1
	if o.lo.xorS != old.lo.xorS {
		return o.lo.query(old.lo, l, m)
	}
	return o.ro.query(old.ro, m+1, r)
}

func cf1771F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	rand.Seed(time.Now().UnixNano())
	const R int = 1e9 + 1

	var n, v, q, l, r, ans int
	Fscan(in, &n)
	t := make([]*node71, n+1)
	t[0] = &node71{lo: null, ro: null}
	mp := make(map[int]int, n)
	for i := 1; i <= n; i++ {
		Fscan(in, &v)
		if _, ok := mp[v]; !ok {
			mp[v] = rand.Int()
		}
		t[i] = t[i-1].modify(0, R, v, mp[v])
	}
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &l, &r)
		l ^= ans
		r ^= ans
		ans = t[r].query(t[l-1], 0, R)
		if ans > 1e9 {
			ans = 0
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1771F(os.Stdin, os.Stdout) }
