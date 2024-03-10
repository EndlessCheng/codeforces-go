package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
var nodes23 [24e6]struct{ lo, ro, sum int32 }
var pid23 int32 = -1

func build23(l, r int) int32 {
	pid23++
	o := pid23
	if l == r {
		return pid23
	}
	m := (l + r) >> 1
	nodes23[o].lo = build23(l, m)
	nodes23[o].ro = build23(m+1, r)
	return o
}

func add23(old int32, l, r, i int) int32 {
	pid23++
	o := pid23
	nodes23[o] = nodes23[old]
	if l == r {
		nodes23[o].sum++
		return o
	}
	m := (l + r) >> 1
	if i <= m {
		nodes23[o].lo = add23(nodes23[o].lo, l, m, i)
	} else {
		nodes23[o].ro = add23(nodes23[o].ro, m+1, r, i)
	}
	nodes23[o].sum = nodes23[nodes23[o].lo].sum + nodes23[nodes23[o].ro].sum
	return o
}

func countRange23(o, old int32, l, r, low, high int) int {
	if high < l || r < low {
		return 0
	}
	if low <= l && r <= high {
		return int(nodes23[o].sum - nodes23[old].sum)
	}
	m := (l + r) >> 1
	if high <= m {
		return countRange23(nodes23[o].lo, nodes23[old].lo, l, m, low, high)
	}
	if m < low {
		return countRange23(nodes23[o].ro, nodes23[old].ro, m+1, r, low, high)
	}
	return countRange23(nodes23[o].lo, nodes23[old].lo, l, m, low, high) + countRange23(nodes23[o].ro, nodes23[old].ro, m+1, r, low, high)
}

func cf323C(_r io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	_i, _n, buf := 0, 0, make([]byte, 1<<12)
	rc := func() byte {
		if _i == _n {
			_n, _ = _r.Read(buf)
			if _n == 0 {
				return 0
			}
			_i = 0
		}
		b := buf[_i]
		_i++
		return b
	}
	r := func() (x int) {
		b := rc()
		for ; '0' > b; b = rc() {
		}
		for ; '0' <= b; b = rc() {
			x = x*10 + int(b&15)
		}
		return
	}

	n := r()
	mp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		mp[r()] = i
	}

	t := make([]int32, n+1)
	t[0] = build23(1, n)
	for i := 1; i <= n; i++ {
		t[i] = add23(t[i-1], 1, n, mp[r()])
	}

	x := -1
	for m := r(); m > 0; m-- {
		l1 := (r()+x)%n + 1
		r1 := (r()+x)%n + 1
		if l1 > r1 {
			l1, r1 = r1, l1
		}
		l2 := (r()+x)%n + 1
		r2 := (r()+x)%n + 1
		if l2 > r2 {
			l2, r2 = r2, l2
		}
		x = countRange23(t[r2], t[l2-1], 1, n, l1, r1)
		Fprintln(out, x)
	}
}

//func main() { cf323C(os.Stdin, os.Stdout) }
