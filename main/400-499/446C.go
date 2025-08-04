package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
const mod46 = 1_000_000_009

var fib46 []int

type pair46 struct{ s, t int }

type seg46 []struct {
	l, r, sum int
	todo      pair46
}

func (t seg46) apply(o int, f pair46) {
	cur := &t[o]
	sz := cur.r - cur.l + 1
	cur.sum = (cur.sum + fib46[sz]*f.s + (fib46[sz+1]-1)*f.t) % mod46
	cur.todo.s = (cur.todo.s + f.s) % mod46
	cur.todo.t = (cur.todo.t + f.t) % mod46
}

func (t seg46) maintain(o int) {
	t[o].sum = (t[o<<1].sum + t[o<<1|1].sum) % mod46 // 这个取模是多余的
}

func (t seg46) spread(o int) {
	f := t[o].todo
	if f == (pair46{}) {
		return
	}
	t.apply(o<<1, f)
	lsz := t[o<<1].r - t[o<<1].l + 1
	// 注意这里会用到 fib[0]
	t.apply(o<<1|1, pair46{(fib46[lsz-1]*f.s + fib46[lsz]*f.t) % mod46, (fib46[lsz]*f.s + fib46[lsz+1]*f.t) % mod46})
	t[o].todo = pair46{}
}

func (t seg46) build(in io.Reader, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		Fscan(in, &t[o].sum)
		return
	}
	m := (l + r) >> 1
	t.build(in, o<<1, l, m)
	t.build(in, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg46) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.apply(o, pair46{fib46[t[o].l-l+1], fib46[t[o].l-l+2]})
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func (t seg46) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func cf446C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, op, l, r int
	Fscan(in, &n, &m)
	fib46 = make([]int, n+2)
	fib46[1] = 1
	for i := 2; i < len(fib46); i++ {
		fib46[i] = (fib46[i-1] + fib46[i-2]) % mod46
	}

	t := make(seg46, 2<<bits.Len(uint(n-1)))
	t.build(in, 1, 0, n-1)

	for range m {
		Fscan(in, &op, &l, &r)
		l--
		r--
		if op == 1 {
			t.update(1, l, r)
		} else {
			Fprintln(out, t.query(1, l, r)%mod46)
		}
	}
}

//func main() { cf446C(bufio.NewReader(os.Stdin), os.Stdout) }
