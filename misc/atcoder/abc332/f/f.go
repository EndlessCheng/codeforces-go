package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
	"os"
)

// https://github.com/EndlessCheng
const mod = 998244353

type pair struct{ mul, add int }

type seg []struct {
	l, r int
	e    int
	todo pair
}

var todoInit = pair{1, 0}

func mergeTodo(f, old pair) pair {
	return pair{f.mul * old.mul % mod, (f.mul*old.add + f.add) % mod}
}

func (t seg) apply(o int, f pair) {
	cur := &t[o]
	cur.e = (f.mul*cur.e + f.add) % mod
	cur.todo = mergeTodo(f, cur.todo)
}

func (t seg) spread(o int) {
	f := t[o].todo
	if f == todoInit {
		return
	}
	t.apply(o<<1, f)
	t.apply(o<<1|1, f)
	t[o].todo = todoInit
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].e = a[l]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
}

func (t seg) update(o, l, r int, f pair) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, f)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, f)
	}
	if m < r {
		t.update(o<<1|1, l, r, f)
	}
}

func (t seg) spreadAll(o int, out io.Writer) {
	if t[o].l == t[o].r {
		Fprint(out, (t[o].e+mod)%mod, " ")
		return
	}
	t.spread(o)
	t.spreadAll(o<<1, out)
	t.spreadAll(o<<1|1, out)
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, x int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	t := make(seg, 2<<bits.Len(uint(n-1)))
	t.build(a, 1, 0, n-1)
	for ; m > 0; m-- {
		Fscan(in, &l, &r, &x)
		p := pow(r-l+1, mod-2)
		t.update(1, l-1, r-1, pair{1 - p, p * x % mod})
	}
	t.spreadAll(1, out)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
